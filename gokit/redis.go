package gokit

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// RedisClusterConnect - returns a connection to a redis cluster
// Expects:-
// REDIS_HOSTS to be a list of hosts separated by a comma i.e HOST1:port,HOST2:port,HOST3:port
func RedisClusterConnect(password string, hosts []string) (client *redis.ClusterClient, err error) {
	Log("%s Connecting to redis cluster...", LogPrefix)
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    hosts,
		Password: password,
	})
	if client == nil {
		return nil, fmt.Errorf("%s unable to connect to redis", LogPrefix)
	}
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("%s unable to connect to redis", LogPrefix)
	}
	if pong == "PONG" {
		Log("%s Connected to redis cluster...", LogPrefix)
	}
	return client, nil
}

// RedisConnect - returns a connection to REDIS_HOST:port
func RedisConnect(username, password, host string, db int) (client *redis.Client, err error) {
	Log("%s Connecting to redis Host::%s", LogPrefix, host)
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Username: username,
		Password: password,
		DB:       db,
	})
	if client == nil {
		return nil, fmt.Errorf("%s unable to connect to redis", LogPrefix)
	}
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("%s unable to connect to redis", LogPrefix)
	}
	if pong == "PONG" {
		Log("%s Connected to redis...", LogPrefix)
	}
	return client, nil
}
