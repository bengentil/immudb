(window.webpackJsonp=window.webpackJsonp||[]).push([[17],{580:function(t,e,r){},585:function(t,e,r){"use strict";r(17),r(9),r(10),r(6),r(5),r(8);var n=r(0),o=(r(580),r(38));function c(object,t){var e=Object.keys(object);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(object);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(object,t).enumerable}))),e.push.apply(e,r)}return e}function l(t){for(var i=1;i<arguments.length;i++){var source=null!=arguments[i]?arguments[i]:{};i%2?c(Object(source),!0).forEach((function(e){Object(n.a)(t,e,source[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(source)):c(Object(source)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(source,e))}))}return t}e.a=o.a.extend({name:"v-divider",props:{inset:Boolean,vertical:Boolean},render:function(t){var e;return this.$attrs.role&&"separator"!==this.$attrs.role||(e=this.vertical?"vertical":"horizontal"),t("hr",{class:l({"v-divider":!0,"v-divider--inset":this.inset,"v-divider--vertical":this.vertical},this.themeClasses),attrs:l({role:"separator","aria-orientation":e},this.$attrs),on:this.$listeners})}})},605:function(t,e,r){},640:function(t,e,r){"use strict";r(605)},653:function(t,e,r){"use strict";r.r(e);var n={name:"UiOrDivider"},o=(r(640),r(26)),c=r(39),l=r.n(c),v=r(570),d=r(585),f=r(571),component=Object(o.a)(n,(function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("section",{staticClass:"or-divider"},[r("v-row",{staticClass:"ma-0 justify-center align-center"},[r("v-col",{staticClass:"pa-0 justify-end align-center",attrs:{cols:"3",sm:"2",md:"2",lg:"1"}},[r("v-divider")],1),t._v(" "),r("v-col",{staticClass:"pa-0 px-4 justify-center align-center",attrs:{cols:"2",sm:"1",md:"1",lg:"1"}},[r("span",{staticClass:"bg overline font-300--text"},[t._v("\n\t\t\t\t"+t._s(t.$t("common.or"))+"\n\t\t\t")])]),t._v(" "),r("v-col",{staticClass:"pa-0 justify-start align-center",attrs:{cols:"3",sm:"2",md:"2",lg:"1"}},[r("v-divider")],1)],1)],1)}),[],!1,null,null,null);e.default=component.exports;l()(component,{VCol:v.a,VDivider:d.a,VRow:f.a})}}]);