package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// AvgSleep how much we sleep in work emulation
const AvgSleep = 50

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/", loadPostsHandle)

	siteHandler := timingMiddleware(siteMux)

	fmt.Println("starting server at :4002")
	http.ListenAndServe(":4002", siteHandler)
}

func emulateWork(ctx context.Context, workName string) {
	defer trackContextTimings(ctx, workName, time.Now())

	rnd := time.Duration(rand.Intn(AvgSleep))
	time.Sleep(time.Millisecond * rnd)
}

func loadPostsHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	emulateWork(ctx, "checkCache")
	emulateWork(ctx, "loadPosts")
	emulateWork(ctx, "loadPosts")
	emulateWork(ctx, "loadPosts")
	time.Sleep(10 * time.Millisecond)
	emulateWork(ctx, "loadSidebar")
	emulateWork(ctx, "loadComments")

	fmt.Fprintln(w, "Request done")
}

func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, timingsKey, &ctxTimings{
			Data: make(map[string]*Timing),
		})
		defer logContextTimings(ctx, r.URL.Path, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func trackContextTimings(ctx context.Context, metricName string, start time.Time) {
	// getting timings from context
	// because of there is an empty interface, we need to cast it to correct data type
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	elapsed := time.Since(start)
	// we lock in case of concurrent records writing into map
	timings.Lock()
	defer timings.Unlock()
	// if there is no metrics yet - we create it, else we will append to existing
	if metric, metricExist := timings.Data[metricName]; !metricExist {
		timings.Data[metricName] = &Timing{
			Count:    1,
			Duration: elapsed,
		}
	} else {
		metric.Count++
		metric.Duration += elapsed
	}
}

type Timing struct {
	Count    int
	Duration time.Duration
}

type ctxTimings struct {
	sync.Mutex
	Data map[string]*Timing
}

// linter will complain if we will use basic types in Value of context
// i.e. this way is more secure to razgranichivat'
type key int

const timingsKey key = 1

func logContextTimings(ctx context.Context, path string, start time.Time) {
	// getting timings from context
	// because of there is an empty interface, we need to cast it to correct data type
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	totalReal := time.Since(start)
	buf := bytes.NewBufferString(path)
	var total time.Duration
	for timing, value := range timings.Data {
		total += value.Duration
		buf.WriteString(fmt.Sprintf("\n\t%s(%d): %s", timing, value.Count, value.Duration))
	}
	buf.WriteString(fmt.Sprintf("\n\ttotal: %s", totalReal))
	buf.WriteString(fmt.Sprintf("\n\ttracked: %s", total))
	buf.WriteString(fmt.Sprintf("\n\tunkn: %s", totalReal-total))

	fmt.Println(buf.String())
}
