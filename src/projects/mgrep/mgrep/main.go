package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

func discoverFiles(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	var fullPath string

	for _, entry := range entries {
		fullPath = filepath.Join(path, entry.Name())
		if entry.IsDir() {
			discoverFiles(wl, fullPath)
		} else {
			wl.Add(worklist.NewJob(fullPath))
		}
	}
}

var args struct {
	SearcTerm string `arg:"positional,required"`
	SearchDir string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	var workersWg sync.WaitGroup

	wl := worklist.New(100)

	numWorkers := 10

	results := make(chan worker.Result, 100)

	workersWg.Add(1)
	go func() {
		defer workersWg.Done()
		discoverFiles(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				job := wl.Next()
				if job.Path != "" {
					workerResults := worker.FindInFile(job.Path, args.SearcTerm)
					if workerResults != nil {
						for _, result := range workerResults.Inner {
							results <- result
						}
					}
				} else {
					return
				}
			}
		}()
	}

	blockWorkerWg := make(chan struct{})

	go func() {
		workersWg.Wait()
		close(blockWorkerWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <-results:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
			case <-blockWorkerWg:
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()

	displayWg.Wait()
}
