package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProtoBuilder(t *testing.T) {
	pb := &protoBuilder{}
	err := pb.buildProtosFromAPI()
	require.NoError(t, err)
}
