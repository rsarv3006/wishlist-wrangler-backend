package coresharedregistry

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hashicorp/consul/api"
)

type ServiceRegistry struct {
	client    *api.Client
	services  map[string][]string
	mu        sync.RWMutex
	serviceId *string
}

func NewServiceRegistry(consulAddr string) (*ServiceRegistry, error) {
	config := api.DefaultConfig()
	config.Address = consulAddr
	consulClient, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ServiceRegistry{
		client:   consulClient,
		services: make(map[string][]string),
	}, nil
}

func (sr *ServiceRegistry) RegisterService(serviceName, host string, port int) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", host, port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	sr.serviceId = &serviceName

	if err := sr.client.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	serviceKey := fmt.Sprintf("%s:%d", host, port)
	sr.services[serviceName] = append(sr.services[serviceName], serviceKey)

	handleShutdown(sr)

	return nil
}

func (sr *ServiceRegistry) DiscoverServices(serviceName string) ([]string, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	services, _, err := sr.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	serviceAddrs := []string{}
	for _, service := range services {
		addr := fmt.Sprintf("%s:%d",
			service.Service.Address,
			service.Service.Port,
		)
		serviceAddrs = append(serviceAddrs, addr)
	}

	return serviceAddrs, nil
}

func (sr *ServiceRegistry) DeregisterService() error {
	return sr.client.Agent().ServiceDeregister(*sr.serviceId)
}

func handleShutdown(sr *ServiceRegistry) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v. Disconnecting from the discovery service.", sig)
		if err := sr.DeregisterService(); err != nil {
			log.Printf("Failed to deregister service: %v", err)
		}
	}()
}
