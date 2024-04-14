import matplotlib.pyplot as plt

goroutines = [0, 1, 10, 100, 1000]

processing_times = [2.0656, 2.0872, 0.5177, 1.0465, 1.5679]

plt.figure(figsize=(10, 6))
plt.plot(goroutines, processing_times, marker='o', linestyle='-')
plt.xlabel('Number of Goroutines')
plt.ylabel('Processing Time (ms)')
plt.grid(True)
plt.xticks(goroutines)
plt.show()
