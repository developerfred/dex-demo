package storeutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementSeq(t *testing.T) {
	ctx, sk := mockApp(t)
	assert.True(t, NewEntityID(1).Equals(IncrementSeq(ctx, sk, testKey)))
	assert.True(t, NewEntityID(2).Equals(IncrementSeq(ctx, sk, testKey)))
}

func TestGetSeq(t *testing.T) {
	ctx, sk := mockApp(t)
	assert.True(t, ZeroEntityID.Equals(GetSeq(ctx, sk, testKey)))
	IncrementSeq(ctx, sk, testKey)
	assert.True(t, NewEntityID(1).Equals(GetSeq(ctx, sk, testKey)))
}
