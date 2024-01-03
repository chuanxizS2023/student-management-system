<template>
  <div>
    <h2>Add Student</h2>
    <form @submit.prevent="addStudent">
      <input type="text" v-model="student.name" placeholder="Name" />
      <input type="date" v-model="student.enrollDate" placeholder="Enroll Date" />
      <button type="submit">Add Student</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      student: {
        name: '',
        enrollDate: '',
      },
    };
  },
  methods: {
    addStudent() {
      // Placeholder for actual API call
      alert(`Adding student: ${this.student.name} with date: ${this.student.enrollDate}`);
      // Reset the form
      this.student.name = '';
      this.student.enrollDate = '';
      // Here you would make an axios call to your backend to add the student
    },
  },
};
</script>
