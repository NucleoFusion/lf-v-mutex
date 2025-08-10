import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

sns.set_theme(style="whitegrid", palette="Set2")
sns.set_context("notebook", font_scale=1.2)

df = pd.read_csv("../results/result.csv")

metrics = ['Throughput', 'Latency', 'Memory', 'CPUUtil', 'GCPauses']

for metric in metrics:
    plt.figure(figsize=(8, 5))
    sns.barplot(
        data=df,
        x='MapImpl',
        y=metric,
        hue='OpType'
    )
    plt.title(f'{metric} by Map Implementation and Operation Type')
    plt.ylabel(metric)
    plt.xlabel('Map Implementation')
    plt.legend(title='OpType')
    if metric == 'Memory':
        plt.yscale('log')  # Log scale for Memory due to huge values
    plt.tight_layout()
    plt.savefig(f'{metric.lower()}_barplot.png', dpi=300)
    plt.show()
