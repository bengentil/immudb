(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{504:function(e,t,l){"use strict";l.r(t);var r={name:"OutputCodeFilter",props:{filter:{type:String,default:""}},data:()=>({value:"",items:["all","stdout","stderr","execerr"]}),watch:{filter:{immediate:!0,handler(e,t){e!==t&&(this.value=e)}},value(e){this.$emit("update:filter",e)}}},n=l(14),d=l(18),c=l.n(d),o=l(558),component=Object(n.a)(r,(function(){var e=this,t=e.$createElement;return(e._self._c||t)("v-select",{staticClass:"ma-0 mt-1 pa-0",class:{"gray--text text--darken-1":!e.$vuetify.theme.dark,"gray--text text--lighten-1":e.$vuetify.theme.dark},attrs:{color:"grey darken-2",items:e.items,label:e.$t("output.code.filter.label"),"hide-details":"",dense:""},model:{value:e.value,callback:function(t){e.value=t},expression:"value"}})}),[],!1,null,null,null);t.default=component.exports;c()(component,{VSelect:o.a})}}]);