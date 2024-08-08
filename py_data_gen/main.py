import pandas as pd
from ctgan import CTGAN
from sqlalchemy import create_engine
from faker import Faker

def pg_conn(host, dbname, user, password, table_name):
    try:
        connection_string = f'postgresql://{user}:{password}@{host}/{dbname}'
        
        engine = create_engine(connection_string)

        query = f"SELECT * FROM {table_name};"
        
        df = pd.read_sql_query(query, engine)
        
        return df

    except Exception as e:
        print(f"Error: {e}")
        return None

# sample model and genreate synthetic data
def generate_synthetic_data(df, epochs=100, batch_size=500, num_samples=1000):

    df_filtered = df.drop(columns=['id'])

    categorical_columns = ['first_name', 'last_name', 'account_number', 'routing_number', 'card_number']
    

    # fit model
    model = CTGAN(epochs=epochs, batch_size=batch_size)
    model.fit(df_filtered, discrete_columns=categorical_columns)
    
    # 
    
    # sample the model and generate synthetic data
    synthetic_data = model.sample(num_samples)
    
    # Generate new UUIDs using Faker
    fake = Faker()
    synthetic_data["id"] = [str(fake.uuid4()) for _ in range(num_samples)]
    
    return synthetic_data

def main():
    host = 'ep-late-cherry-a5k4dfkr.us-east-2.aws.neon.tech'
    dbname = 'syndata'
    user = 'evis'
    password = 'lY7JyhzuLDT5'
    table_name = 'users'

    df = pg_conn(host, dbname, user, password, table_name)

    if df is not None:
        print("Data loaded successfully.")
        print("Starting to generate synthetic data...")
        
        synthetic_data = generate_synthetic_data(df)
        
        print("Synthetic data generated successfully.")
        print(synthetic_data.head())

        # save to a CSV file
        synthetic_data.to_csv('synthetic_users.csv', index=False)
        print("Synthetic data saved to 'synthetic_users.csv'.")

if __name__ == "__main__":
    main()
