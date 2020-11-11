import mongoose from 'mongoose';

mongoose.connect('mongodb://localhost:27017/', { useNewUrlParser: true, useUnifiedTopology: true, dbName: 'boilerplate' });
const db = mongoose.connection;

db.on('error', e => console.error('MongoDB connection error:', e));
db.on('connected', () => {
  console.log('Connection to MongoDB established');
});
db.on('disconnected', () => {
  console.warn('Lost connection to MongoDB');
});
db.on('reconnect', () => {
  console.log('Reconnected to MongoDB');
});
db.on('close', () => {
  console.warn('Closed connection to MongoDB');
});
