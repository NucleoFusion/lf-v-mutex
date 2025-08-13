import pandas as pd
import json

def load_data(path):
    with open(path,"r") as f:
        data = json.load(f)

    records = []
    for i,entry in enumerate(data,1):
        for key in ['readHeavy','writeHeavy','balanced']:
            row = entry[key].copy()
            op_type = row.pop('OpType')
            map_impl = row['MapImpl']
            records.append(((map_impl, op_type,i), row))

    index = pd.MultiIndex.from_tuples([r[0] for r in records], names=['MapImpl', 'OpType', 'Trial'])
    df = pd.DataFrame([r[1] for r in records], index=index)


    avg = df.groupby('OpType').mean(numeric_only=True)
    meta = df[['Scenario']].groupby(['MapImpl','OpType']).first()
    
    res = avg.join(meta)
    
    return res 

df_lockfree = load_data("../results/lockfree.json")
df_rwmutex = load_data("../results/rwmutex.json")
# df_atomic = load_data("../results/atomicmap.json")
df_partial = load_data("../results/partiallockfree.json")
#
# print(df_lockfree)

result = pd.concat([df_rwmutex, df_lockfree, df_partial], ignore_index=False)

res = result.columns.tolist()               # get list of columns
res.insert(0, res.pop(res.index('Scenario')))  # remove 'Scenario' from current position and insert at front
result = result[res]

result.to_csv("../results/result.csv")

