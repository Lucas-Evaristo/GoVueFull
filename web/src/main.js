const app = Vue.createApp({
  data() {
    return {
      courseGoal: "stringy",
      vueLink: "https://vuejs.org",
    };
  },
  methods: {
    outputGoal() {
      const randomNumber = Math.random();
      if (randomNumber < 0.5) {
        return "Learn Vue";
      } else {
         return 'Master Vue'
      }
    },
  },
});

const app2 = Vue.createApp({
   data(){
      return{
         courseGoal2: 'Another stringy'
      }
   }
})

app.mount("#user-goal");
