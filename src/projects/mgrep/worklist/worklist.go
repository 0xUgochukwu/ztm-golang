package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (wl *Worklist) Add(work Entry) {
	wl.jobs <- work
}

func (wl *Worklist) Next() Entry {
	job := <-wl.jobs
	return job
}

func New(buffSize int) Worklist {
	return Worklist{make(chan Entry, buffSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (wl *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		wl.Add(NewJob(""))
	}
}
