<template>
  <div class="row">
    <div class="row no-gutters">
      <div class="container">
        <label for="group-name">グループ名:</label>
        <input
          type="text"
          class="group-name"
          placeholder="新しいグループ名を入力"
          style="boder:none"
        />
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style scoped>
.container {
  background-color: white;
  height: 5vh;
  width: 100%;
}

.group-name {
  width: 70vw;
  outline: 0;
}
</style>
