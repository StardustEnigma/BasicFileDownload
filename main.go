package main

import (
	"fmt"
	"time"
)

func (d DownloadJob) Display(){
	fmt.Println(d.FileName)
}

type DownloadJob struct {
	ID int
	FileName string
	SizeMB int
	Status string
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

func Download(job DownloadJob){
	fmt.Println("Starting",job.FileName)

	time.Sleep(2*time.Second)

	fmt.Println("Downloaded",job.FileName)
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
	stats["completed"]++
	fmt.Println(stats)

	logger("Download Started")
	PrintAll(jobs)
	
	for _ ,job := range jobs{
		go Download(job)
	}
	time.Sleep(3*time.Second)
	
}