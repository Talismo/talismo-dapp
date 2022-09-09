<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>Profile Search</h1>
      </v-col>
      <v-col cols="12">
        <hello-world />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HelloWorld from "../../components/HelloWorld.vue";

export default {
  name: "Home",

  components: {
    HelloWorld,
  },
};
</script>
