// +build streams
/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"context"
	"fmt"
	ic "github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/codenotary/immudb/pkg/server/servertest"
	"github.com/codenotary/immudb/pkg/server/sessions"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

func TestSession_OpenCloseSession(t *testing.T) {
	options := server.DefaultOptions()
	bs := servertest.NewBufconnServer(options)

	defer os.RemoveAll(options.Dir)
	defer os.Remove(".state-")

	bs.Start()
	defer bs.Stop()

	client := ic.DefaultClient().WithOptions(ic.DefaultOptions().WithDialOptions([]grpc.DialOption{grpc.WithContextDialer(bs.Dialer), grpc.WithInsecure()}))

	serverUUID, sessionID, err := client.OpenSession(context.TODO(), []byte(`immudb`), []byte(`immudb`), "defaultdb")
	require.NoError(t, err)
	require.NotNil(t, serverUUID)
	require.NotNil(t, sessionID)

	client.Set(context.TODO(), []byte("my"), []byte("session"))

	err = client.CloseSession(context.TODO())
	require.NoError(t, err)
}

func TestSession_OpenCloseSessionMulti(t *testing.T) {
	sessOptions := &sessions.Options{
		SessionGuardCheckInterval: time.Millisecond * 100,
		MaxSessionIdle:            time.Millisecond * 2000,
		MaxSessionAge:             time.Millisecond * 4000,
		Timeout:                   time.Millisecond * 2000,
	}
	options := server.DefaultOptions().WithSessionOptions(sessOptions)
	bs := servertest.NewBufconnServer(options)

	defer os.RemoveAll(options.Dir)
	defer os.Remove(".state-")

	bs.Start()
	defer bs.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			client := ic.DefaultClient().WithOptions(ic.DefaultOptions().
				WithDialOptions([]grpc.DialOption{grpc.WithContextDialer(bs.Dialer), grpc.WithInsecure()}).
				WithHeartBeatFrequency(time.Millisecond * 100))
			serverUUID, sessionID, err := client.OpenSession(context.TODO(), []byte(`immudb`), []byte(`immudb`), "defaultdb")
			require.NoError(t, err)
			require.NotNil(t, serverUUID)
			require.NotNil(t, sessionID)

			min := 10
			max := 100
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(max-min)+min))

			client.Set(context.TODO(), []byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", i)))

			err = client.CloseSession(context.TODO())
			require.NoError(t, err)
			wg.Done()
		}(i)
	}
	wg.Wait()
	require.Equal(t, 0, bs.Server.Srv.SessManager.CountSession())
}
