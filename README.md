# containertools

Tools for happy existence in a container

The containertools project is intended to yield a fairly generic library for detecting that a process is running in a container, and allow us to take appropriate action based on that determination.

The ultimate goal is to be able to provide a set of functions that allow us to raise some kind of semaphore to our containerization engine (without needing access to the docker socket or to any sort of privilege), so that it can take action on behalf of the container (for instance, committing changed content, or restarting after shutdown).
