<template>
    <div>
      <h2>Teacher List</h2>
      <ul>
        <li v-for="teacher in teachers" :key="teacher.id">
          {{ teacher.name }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        teachers: [
          // Populate with data from backend
        ],
      };
    },
    // Fetch teachers data from backend
  };
  </script>
  