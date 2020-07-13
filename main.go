package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	"github.com/docker/distribution/reference"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) != 2 {
		return errors.New("please provide a reference")
	}

	ref, err := reference.ParseNormalizedNamed(args[1])
	if err != nil {
		return err
	}

	err = fetch(context.Background(), ref, createResolver())
	if err != nil {
		return err
	}
	return nil
}

func createResolver() remotes.Resolver {

	return docker.NewResolver(docker.ResolverOptions{
		Authorizer: docker.NewDockerAuthorizer(),
		PlainHTTP:  false,
	})
}

func fetch(ctx context.Context, ref reference.Named, resolver remotes.Resolver) error {
	_, indexDescriptor, err := resolver.Resolve(ctx, ref.String())
	if err != nil {
		return err
	}

	fetcher, err := resolver.Fetcher(ctx, ref.String())
	if err != nil {
		return err
	}
	reader, err := fetcher.Fetch(ctx, indexDescriptor)
	if err != nil {
		return err
	}
	defer reader.Close()

	result, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	var index ocischemav1.Index
	if err := json.Unmarshal(result, &index); err != nil {
		return err
	}

	b, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
