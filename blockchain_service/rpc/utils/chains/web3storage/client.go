package web3storage

import (
	"context"
	"fmt"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/config"
	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"io/fs"
	"os"
)

type W3sClient struct {
	Client w3s.Client
}

func (w *W3sClient) InitClient(ctx context.Context, config config.Config) (err error) {
	w.Client, err = w3s.NewClient(
		w3s.WithEndpoint(config.Chains.Web3StorageConfig.Endpoint),
		w3s.WithToken(config.Chains.Web3StorageConfig.Token),
	)
	if err != nil {
		return
	}

	return
}

func (w *W3sClient) Put(context context.Context, file *os.File) (cid string, err error) {
	c, err := w.Client.Put(context, file)
	if err != nil {
		return "", err
	}
	return c.String(), nil
}

func (w *W3sClient) Get(context context.Context, cidString string) (file fs.File, err error) {
	c, err := cid.Parse(cidString)
	if err != nil {
		return nil, err
	}
	resp, err := w.Client.Get(context, c)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got status %d, wanted %d", resp.StatusCode, 200)
	}

	f, fsys, err := resp.Files()
	if err != nil {
		return nil, fmt.Errorf("failed to read files: %v", err)
	}

	// List directory entries
	if d, ok := f.(fs.ReadDirFile); ok {
		ents, _ := d.ReadDir(0)
		for _, ent := range ents {
			fmt.Println(ent.Name())
		}
	}

	// Walk whole directory contents (including nested directories)
	fs.WalkDir(fsys, "/", func(path string, d fs.DirEntry, err error) error {
		info, _ := d.Info()
		fmt.Printf("%s (%d bytes)\n", path, info.Size())
		return err
	})

	// Open a file in a directory
	file, err = os.Open("D:\\workspace\\golang_workspace\\src\\github.com\\JopenChen\\star_guard\\blockchain_service\\rpc\\d41d8cd98f00b204e9800998ecf8427e")

	return
}
