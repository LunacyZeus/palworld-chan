import{d as V,r as a,O as D,y as F,b as l,C as o,D as E,P as B,H as k}from"./9a74f89d.js";import{C as y}from"./bd857e79.js";import"./74dee749.js";import{F as A}from"./6c55285d.js";import{_ as C}from"./f155ca5d.js";import{F as x}from"./9d312194.js";import{_ as S}from"./0ee26181.js";import"./8f51a641.js";import"./e27e03ce.js";import"./709c511c.js";const U={style:{margin:"16px"}},T=V({__name:"SaveBackupSetting",setup($){const n=a("/home/docker-palworld-dedicated-server/game/Pal/Saved"),s=a("/root/backup/"),d=a(1800),i=a(200),m=a(""),p=a(""),c=a(""),v=r=>`${r} \u4E0D\u5408\u6CD5\uFF0C\u8BF7\u91CD\u65B0\u8F93\u5165`,g=r=>{};return(r,e)=>{const t=A,_=y,b=D,f=x;return E(),F("div",null,[l(C),l(f,{onSubmit:g},{default:o(()=>[l(_,{inset:"",style:{margin:"6px"},title:"\u672C\u5730\u5907\u4EFD"},{default:o(()=>[l(t,{label:"\u5B58\u6863\u8DEF\u5F84","input-align":"right",center:!0,border:!1,modelValue:n.value,"onUpdate:modelValue":e[0]||(e[0]=u=>n.value=u)},null,8,["modelValue"]),l(t,{label:"\u5907\u4EFD\u8DEF\u5F84","input-align":"right",center:!0,border:!1,modelValue:s.value,"onUpdate:modelValue":e[1]||(e[1]=u=>s.value=u)},null,8,["modelValue"]),l(t,{label:"\u5B58\u6863\u5468\u671F(\u79D2)","input-align":"right",center:!0,border:!1,modelValue:d.value,"onUpdate:modelValue":e[2]||(e[2]=u=>d.value=u),rules:[{required:!0,message:"\u5B58\u6863\u5468\u671F\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:v}]},null,8,["modelValue","rules"]),l(t,{label:"\u6EDA\u52A8\u4FDD\u5B58\u6570","input-align":"right",center:!0,border:!1,modelValue:i.value,"onUpdate:modelValue":e[3]||(e[3]=u=>i.value=u),rules:[{required:!0,message:"\u6EDA\u52A8\u4FDD\u5B58\u6570\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:v}]},null,8,["modelValue","rules"])]),_:1}),l(_,{inset:"",style:{margin:"6px"},title:"\u4E03\u725B\u5907\u4EFD"},{default:o(()=>[l(t,{label:"accessToken","input-align":"right",center:!0,border:!1,modelValue:m.value,"onUpdate:modelValue":e[4]||(e[4]=u=>m.value=u),placeholder:"\u4E03\u725BaccessToken"},null,8,["modelValue"]),l(t,{label:"secretKey","input-align":"right",center:!0,border:!1,modelValue:p.value,"onUpdate:modelValue":e[5]||(e[5]=u=>p.value=u),placeholder:"\u4E03\u725BsecretKey"},null,8,["modelValue"]),l(t,{label:"bucket","input-align":"right",center:!0,border:!1,modelValue:c.value,"onUpdate:modelValue":e[6]||(e[6]=u=>c.value=u),placeholder:"\u4E03\u725Bbucket"},null,8,["modelValue"])]),_:1}),B("div",U,[l(b,{round:"",block:"",type:"primary","native-type":"submit"},{default:o(()=>[k(" \u4FEE\u6539 ")]),_:1})])]),_:1})])}}});const h=S(T,[["__scopeId","data-v-0d91bfea"]]);export{h as default};
