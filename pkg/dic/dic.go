package dic

import "context"

type container struct {
	ctx          context.Context
	dependencies map[string]interface{}
}

type Container interface {
	Register(name string, lazyLoader LazyLoadingFunc)
	Get(name string) interface{}
	Context() context.Context
}

// LazyLoadingFunc is a function that returns a dependency
type LazyLoadingFunc func(c Container) interface{}

// New creates a new container
func New() Container {
	return &container{
		dependencies: make(map[string]interface{}),
	}
}

// Register registers a dependency
func (c *container) Register(name string, lazyLoader LazyLoadingFunc) {
	c.dependencies[name] = lazyLoader
}

// Get gets a dependency
func (c *container) Get(name string) interface{} {
	if _, ok := c.dependencies[name]; !ok {
		return nil
	}

	if _, ok := c.dependencies[name].(LazyLoadingFunc); ok {
		c.dependencies[name] = c.dependencies[name].(LazyLoadingFunc)(c)
	}

	return c.dependencies[name]
}

// Context gets the context
func (c *container) Context() context.Context {
	return c.ctx
}
