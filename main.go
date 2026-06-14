package main

import (
	"fmt"
	"time"
	"sync" 
)

type DownloadJob struct {
	ID int
	FileName string
	SizeMB int
	Status string
}
func (d DownloadJob) Display(){
	fmt.Println(d.FileName)
}
func logger(msg string){
	fmt.Println(msg)
}

func IdGenerator() func() int{
	id :=0
	return func() int {
		id++
		return id
	}
} 

type Processor interface{
	Process()
}

func (d DownloadJob) Process(){
	fmt.Println("Processing",d.FileName)
}

func PrintAll[T any](items []T){
	for _,item :=range items{
		fmt.Println(item)
	}
}

func Download(job DownloadJob, successChan chan string,failureChan chan string,mu *sync.Mutex,stats map[string]int){
	fmt.Println("Starting",job.FileName)
	defer fmt.Println("Finished",job.FileName)

	time.Sleep(2*time.Second)
	if job.SizeMB > 50 {
		mu.Lock()
		stats["failed"]++
		mu.Unlock()
		
		failureChan <- job.FileName
		return
	}
	mu.Lock()
	stats["completed"]++
	mu.Unlock()
	successChan <- job.FileName
	
}


func main(){
	jobs := []DownloadJob{}
	nextId :=IdGenerator()
	jobs = append(jobs,
		DownloadJob{nextId(),"movie.mp4",100,"pending"},
		DownloadJob{nextId(), "song.mp3", 20, "pending"},
		DownloadJob{nextId(), "book.pdf", 5, "pending"},
	)
	fmt.Println(len(jobs))
	fmt.Println(cap(jobs))

	for _,job := range jobs{
		job.Display()
		job.Process()
	}

	stats := map[string]int{
		"completed" : 0,
		"failed":0,
	}
	mu := sync.Mutex{}
	fmt.Println(stats)

	logger("Download Started")
	PrintAll(jobs)

	var p Processor

	for _,job := range jobs{
		p=job
		p.Process()
	}
	successChan := make(chan string,3)
	failureChan := make(chan string,3)
	for _ ,job := range jobs{
		go Download(job,successChan,failureChan,&mu,stats)
		
	}
	for i := 0; i < len(jobs); i++ {
		select{
		case file := <- successChan :
			fmt.Println("Success :",file)

		case file := <- failureChan :
			fmt.Println("Failed :",file)
		}
	
	}
	fmt.Println("Final Stats:")
	fmt.Println(stats)
	
	
}