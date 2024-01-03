<template>
  <div id="app">
    <h1>Student Management System</h1>
    <button @click="currentView = 'add-student'">Add Student</button>
    <button @click="currentView = 'add-teacher'">Add Teacher</button>
    <button @click="currentView = 'add-class'">Add Class</button>
    <button @click="currentView = 'manage-grade'">Manage Grade</button>

    <component :is="currentView"></component>

    <student-list></student-list>
    <teacher-list></teacher-list>
    <class-list></class-list>
    <grade-list></grade-list>
  </div>
</template>

<script>
import AddStudent from './components/AddStudent.vue';
import AddTeacher from './components/AddTeacher.vue';
import AddClass from './components/AddClass.vue';
import ManageGrade from './components/ManageGrade.vue';
import StudentList from './components/StudentList.vue';
import TeacherList from './components/TeacherList.vue';
import ClassList from './components/ClassList.vue';
import GradeList from './components/GradeList.vue';

export default {
  name: 'App',
  components: {
    AddStudent,
    AddTeacher,
    AddClass,
    ManageGrade,
    StudentList,
    TeacherList,
    ClassList,
    GradeList,
  },
  data() {
    return {
      currentView: 'add-student',
    };
  },
};
</script>
