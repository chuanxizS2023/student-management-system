<template>
    <div>
      <h2>Student List</h2>
      <ul>
        <li v-for="student in students" :key="student.id">
          {{ student.name }} - Enrolled on: {{ student.enrollDate }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        students: [
          { id: 1, name: 'Alice', enrollDate: '2021-09-01' },
          // ... other students
        ],
      };
    },
  };
  </script>
  