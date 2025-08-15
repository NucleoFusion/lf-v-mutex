# Comparitive Analysis of Concurrency Control Methods



## Abstract

Concurrency is a fundamental concern for high-performance systems, and through the years, different methods have emerged that govern how a shared piece of data can be accessed. In this study, I analyze different concurrency control methods to understand their characteristics and impact on performance. We compare 3 different methods, Lock-Free, RWMutex, & partially Lock-Free. These implementations are further tested on different workloads, such as Read-Heavy, Write-Heavy and Balanced. Several metrics are taken into account, including throughput, latency, CPU utilization, memory utilization, and more.

## Research Paper

## Usage

1. Clone the repo
```
git clone https://github.com/NucleoFusion/lf-v-mutex
```

2. Get into the folder
```
cd lf-v-mutex
```

3. Run the program (Takes a long time, default is 1000*10000)
```
go run .
```

4. (Optionally) change the values for `numRoutines` and `opsPerRoutine` in `run.go` file. Default is 1000 and 10000 respectively.

## Contact

**Author: Lakshit Singh**

Email: lakshit.singh.mail@gmail.com

ORC-ID: 0009-0008-7072-0093
