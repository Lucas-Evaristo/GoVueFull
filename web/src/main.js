
const app = Vue.createApp({
  data() {
    return {
      textoPrueba: "Texto extraido de data()",
      vueLink: "https://vuejs.org",
      counter: 0,
      valor: 5,
      name: "",
    };
  },
  methods: {
    reduceCounter() {
      if (this.counter >= this.valor) {

        this.counter -= this.valor;
      }else{
        alert("sos tarado? ya esta en 0 el contador hijo de la chingada")
      }
    },
    addCounter() {
      this.counter += this.valor;
    },
    submitForm(){
      alert("se envio el formulario")
    }
  },
  computed: {
    fullName(){
      if(this.name == ""){
        return "";
      }else{
        return this.name + " " + "Evaristo";
      }
    }
  },
});

app.mount("#events");

