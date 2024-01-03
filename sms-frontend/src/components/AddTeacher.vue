<template>
    <div>
      <h2>Add Teacher</h2>
      <form @submit.prevent="addTeacher">
        <input type="text" v-model="teacher.name" placeholder="Teacher's Name" />
        <button type="submit">Add Teacher</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        teacher: {
          name: '',
        },
      };
    },
    methods: {
      addTeacher() {
        // Implement the logic to add a teacher
      },
    },
  };
  </script>
  