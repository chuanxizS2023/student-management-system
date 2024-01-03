<template>
    <div>
      <h2>Grade List</h2>
      <ul>
        <li v-for="grade in grades" :key="grade.studentId">
          <!-- Grade details -->
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        grades: [
          // Populate with data from backend
        ],
      };
    },
    // Fetch grades data from backend
  };
  </script>
  