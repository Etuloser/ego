package mysql

import "time"

// Option -.
type Option func(*Client)

// ConnMaxLifetime -.
func ConnMaxLifetime(timeout time.Duration) Option {
	return func(c *Client) {
		c.ConnMaxLifetime = timeout
	}
}

// MaxOpenConns -.
func MaxOpenConns(conns int) Option {
	return func(c *Client) {
		c.MaxOpenConns = conns
	}
}

// MaxIdleConns -.
func MaxIdleConns(conns int) Option {
	return func(c *Client) {
		c.MaxIdleConns = conns
	}
}
