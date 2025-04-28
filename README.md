# 🐳 D-CRDT: Platform-Independent CRDT System with Docker and Kubernetes

**D-CRDT** is a **platform-independent, containerized CRDT (Conflict-Free Replicated Data Type) system** that enables seamless deployment across diverse operating systems and scales dynamically using **Kubernetes orchestration**.

## Overview

Conflict-Free Replicated Data Types (CRDTs) offer strong eventual consistency in distributed systems, but deploying CRDT systems across heterogeneous environments can be challenging.

**D-CRDT** solves this by:
- **Packaging CRDT services into Docker containers** for easy cross-platform deployment.
- **Leveraging Kubernetes** to enable **scalable**, **load-balanced**, and **fault-tolerant** replication.
- Providing **modular CRDT components** (counters, sets, maps) that can be deployed independently or together.

## Key Features

- 📦 **Containerized Deployment**: Build once, run anywhere with Docker.
- ⚖️ **Dynamic Scalability**: Kubernetes manages scaling and load balancing automatically.
- 🔗 **Cross-Platform Compatibility**: Deployable on any OS with Docker/Kubernetes support.
- 🔄 **Flexible CRDT Modules**: Include or exclude different CRDT types as needed.
- 🚀 **Easy Integration**: Lightweight and fast deployment process.

## Technologies Used

- **Go** — CRDT logic and API services.
- **Docker** — containerization of CRDT nodes.
- **Kubernetes** — orchestration for scaling, load balancing, and high availability.

## Motivation

Building distributed systems that are both **consistent** and **easily deployable** across platforms is a major engineering challenge.

**D-CRDT** addresses this by providing:
- A **plug-and-play CRDT system**,
- **No system-specific configuration requirements**,
- **Out-of-the-box scalability** through Kubernetes clusters.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/surzoprovakar/CS6704-Project-D_CRDT.git
   cd CS6704-Project-D_CRDT
