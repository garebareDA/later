<template>
  <div>
    <v-content>
      <v-container>
        <v-tabs centered v-model="tabs">
          <v-tab href="#edit">記事を書く</v-tab>
          <v-tab href="#item">書いた記事</v-tab>
          <v-tab href="#draft">下書き</v-tab>
          <v-tab href="#like">いいねした記事</v-tab>
        </v-tabs>

        <v-tabs-items v-model="tabs">
          <v-tab-item value="edit">
            <editor :titles="title" :texts="text" :uuids="uuid" @drafts="updateDrafts" @publics="updatePublic"/>
          </v-tab-item>

          <v-tab-item value="item">
            <v-card flat>
              <drafts @change="changeToEdit" @reload="itemRelaod"  v-if="reload.item == true" :parts="'item'"/>
            </v-card>
          </v-tab-item>

          <v-tab-item value="draft" >
            <v-card flat>
              <drafts @change="changeToEdit" @reload="draftRelaod" v-if="reload.draft == true" :parts="'draft'"/>
            </v-card>
          </v-tab-item>

          <v-tab-item value="like">
            <v-card flat>
              <drafts @change="changeToEdit" @reload="likeReload" v-if="reload.like == true" :parts="'like'"/>
            </v-card>
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import editor from "../parts/edit.vue";
import drafts from "../parts/drafts.vue";
export default Vue.extend({
  data: () => {
    return {
      title: "",
      text: "",
      uuid: null,
      tabs: "edit",
      update:{
        draft:false,
        item:false,
      },
      reload:{
        draft:true,
        item:true,
        like:true,
      }
    };
  },

  methods: {
    changeToEdit: function(uuid: string, title: string, content: string) {
      this.$data.title = title;
      this.$data.text = content;
      this.$data.uuid = uuid;
      this.$data.tabs = "edit";
    },

    draftRelaod:async function(){
      this.$data.reload.draft = false;
      await this.$nextTick();
      this.$data.reload.draft = true;
    },

    itemRelaod:async function(){
      this.$data.reload.item = false;
      await this.$nextTick();
      this.$data.reload.item = true;
    },

    likeReload:async function() {
      this.$data.reload.like = false;
      await this.$nextTick();
      this.$data.reload.like = true;
    },

    updateDrafts:function(){
      this.$data.update.draft = true;
    },

    updatePublic:function(){
      this.$data.update.draft = true;
      this.$data.update.item = true;
    }
  },

  watch:{
    tabs:function() {
      const draftIsUpdate = this.$data.update.draft;
      if (this.$data.tabs === "draft" && draftIsUpdate){
        this.draftRelaod();
        this.$data.update.draft = false;
      }

      const itemIsUpdate = this.$data.update.item;
      if(this.$data.tabs === "item" && itemIsUpdate){
        this.itemRelaod();
        this.$data.update.item = false;
      }


      const likeIsUpdate = this.$data.update.like;
      if(this.$data.tabs === "like" && likeIsUpdate){
        this.likeReload();
        this.$data.update.like = false;
      }
    }
  },

  components: {
    editor,
    drafts
  }
});
</script>