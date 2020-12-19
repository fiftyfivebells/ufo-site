import numpy as np
import pandas as pd
from sklearn.cluster import KMeans
from sklearn.preprocessing import scale
from sqlalchemy import create_engine

# Connect to the database to get the data
engine = create_engine('postgresql://stephen:stephen@localhost:5432/stephendb')
conn = engine.connect()
query = "SELECT season, latitude, longitude FROM sightings WHERE sighted = 1"

data = pd.read_sql(query, conn)
df = pd.DataFrame(data, columns=['season', 'latitude', 'longitude'])

data = []

# Function for turning a season into a float
def season_to_float(season):
    if season == 'winter':
        return 0.0
    if season == 'spring':
        return 1.0
    if season == 'summer':
        return 2.0
    if season == 'fall':
        return 3.0

    return -1.0

# Make an array of lat, long, and season (as a float)
for index, row in df.iterrows():
    lat = row['latitude']
    lon = row['longitude']
    sea = row['season']
    data.append([float(lon), float(lat), season_to_float(sea)])
        
# Turn the array into a numpy array
data = np.array(data)
        
# Run the kmeans clustering algorithm looking for 6 clusters
model = KMeans(n_clusters=6).fit(scale(data))

a = []

# Store the data in a numpy array as triples
for x in range(len(data[:, 0])):
    a.append([model.labels_[x], data[:, 0][x], data[:, 1][x]])

data_points = np.asarray(a)

data_points = pd.DataFrame(data = data_points,
                           columns = ['label', 'lat', 'long'])

# Save the data to a .csv file for processing
data_points.to_csv('ui/static/clustering.csv', sep=",")
