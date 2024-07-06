<template>
    <div>
      <h2>User Records</h2>
      <ul v-if="records.length">
        <li v-for="record in records" :key="record.ID">
          <p>{{ record.Url }} - {{ record.Type }} - {{ record.Time }} - {{ record.PatientName }}</p>
        </li>
      </ul>
      <p v-else>{{ message }}</p>
    </div>
  </template>
  
  <script>
  import auth from '../services/auth';
  
  export default {
    data() {
      return {
        records: [],
        message: '',
      };
    },
    created() {
      auth.getUserRecords().then(response => {
        this.records = response.records;
      }).catch(error => {
        this.message = error.response.data.error_message;
      });
    },
  };
  </script>
  