import pandas as pd
import sys
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import classification_report
from sqlalchemy import create_engine
import random

engine = create_engine('postgresql://stephen:stephen@localhost:5432/stephendb')
conn = engine.connect()
query1 = "SELECT season, latitude, longitude, sighted FROM sightings"
query2 = "SELECT season, latitude, longitude, sighted FROM sightings ORDER BY random() limit 65000"

# Return a random season
def add_season(season):
    season = ['fall', 'winter', 'spring', 'summer']
    return season[random.randrange(0, 4)]

data = pd.read_sql(query1, conn)
df1 = pd.DataFrame(data, columns=['season', 'latitude', 'longitude', 'sighted'])

data = pd.read_sql(query2, conn)
df2 = pd.DataFrame(data, columns=['season', 'latitude', 'longitude', 'sighted'])

df2['sighted'] = 0

# Randomize the seasons for half of the data to simulate not having
# seen a UFO at that place and time
df2['season'] = df2['season'].apply(add_season)

df = pd.concat([df1, df2])

def convert_season_to_float(season):
    if season == "winter":
        return 0.0
    if season == "spring":
        return 1.0
    if season == "summer":
        return 2.0
    if season == "fall":
        return 3.0

    return -1.0

y = df.iloc[:, 3].values

# Drop columns not being used
x = df.drop(columns=['sighted'])

# Convert season column to numbers
x['season'] = x['season'].apply(convert_season_to_float)

logmodel = LogisticRegression(max_iter=1000)
logmodel.fit(x, y)
logmodel.score(x, y)
prediction = logmodel.predict(x)

# print(classification_report(y, prediction))


# Prediction for a new location and season
def get_prediction(lat, lon, season):
    x = pd.DataFrame([[lat, lon, convert_season_to_float(season)]], columns=['latitude', 'longitude', 'season'])

    probability = logmodel.predict_proba(x)[:, 1]

    return probability[0]


p = get_prediction(sys.argv[1], sys.argv[2], sys.argv[3])

sys.stdout.write(str(p))
sys.exit(0)
