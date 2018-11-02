package cidr_test

import (
	"net"
	"testing"

	"github.com/DavidCai1993/cidr-cli/cidr"
	"github.com/stretchr/testify/require"
)

func TestNextIP(t *testing.T) {
	require := require.New(t)

	ip := net.ParseIP("127.0.0.1")

	require.NotEmpty(ip)

	nextIP := cidr.NextIP(&ip, 1)

	require.Equal("127.0.0.2", nextIP.String())
}
