db = db.getSiblingDB('app_db');

// Create a user specifically for the app_db database
db.createUser({
  user: 'admin',
  pwd: 'admin',
  roles: [
    {
      role: 'readWrite',
      db: 'app_db',
    },
    {
      role: 'dbAdmin',
      db: 'app_db',
    },
  ],
});

// Create initial collection if needed
db.createCollection('initial');

print('Database app_db initialized with user admin');
