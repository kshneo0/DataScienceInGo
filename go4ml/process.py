import pandas as pd

df = pd.read_csv("data/hcvdat0.csv")

df.dtypes

df['Sex'].unique()

gender_dict = {'f':0,'m':1}

df['Sex'].map(gender_dict)

df['Sex'] = df['Sex'].map(gender_dict)

df.head()

df['Category'].value_counts()

df['Category'].unique()

hcdict = {'0=Blood Donor':0, 
    '0s=suspect Blood Donor':1, 
    '1=Hepatitis':1,
    '2=Fibrosis':1, 
    '3=Cirrhosis':1}

df['Category']

df['Target'] = df['Category'].map(hcdict)

df.isna().sum()

df.fillna(0, inplace=True)

df.isna().sum()

df.dtypes

df.head()

df.shape

df.columns

df = df[['Age', 'Sex', 'ALB', 'ALP', 'ALT', 'AST',
       'BIL', 'CHE', 'CHOL', 'CREA', 'GGT', 'PROT', 'Target']]

df

df.head()

train = df.sample(frac=0.7, random_state=200)

test = df.drop(train.index)

train.head()

test.head()

train.to_csv("data/trainhcvData.csv", index=False, header=False)
test.to_csv("data/testhcvData.csv", index=False, header=False)
df.to_csv("data/cleanhcvData.csv")