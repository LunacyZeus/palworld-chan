var c=(_,n,t)=>new Promise((r,e)=>{var l=o=>{try{u(t.next(o))}catch(a){e(a)}},i=o=>{try{u(t.throw(o))}catch(a){e(a)}},u=o=>o.done?r(o.value):Promise.resolve(o.value).then(l,i);u((t=t.apply(_,n)).next())});import{d as F,r as d,O as b,y,b as s,C as p,_ as D,b7 as I,a0 as V,Z as E,D as x,P as B,H as S}from"./e1cebf43.js";import{C}from"./ff51593b.js";import"./926c5ea7.js";import{F as T}from"./0286c5d0.js";/* empty css        */import{_ as w}from"./cff6a0f4.js";import{u as A}from"./9bc92a61.js";import{F as h}from"./394369bc.js";import{_ as N}from"./0ee26181.js";import"./e28c11a8.js";import"./1484675b.js";import"./cdbaf723.js";import"./bbfea913.js";const k={style:{margin:"16px"}},G=F({__name:"Info",setup(_){const n=d(""),t=d(""),r=d(!1),e=A(),l=()=>c(this,null,function*(){yield e.Info(),I(),r.value=!1,e.getServerInfo?(""+e.getResult,V("\u83B7\u53D6\u6210\u529F: "+e.getServerInfo.name),n.value=e.getServerInfo.name,t.value=e.getServerInfo.version):e.getError&&(console.error("Failed to Get Info:",e.getError),E("\u91CD\u542F\u5931\u8D25: "+e.getError))}),i=u=>{r.value=!0,D({duration:0,message:"\u83B7\u53D6\u670D\u52A1\u4FE1\u606F\u4E2D...",loadingType:"spinner"}),l()};return(u,o)=>{const a=T,f=b,v=C,g=h;return x(),y("div",null,[s(w),s(g,{onSubmit:i},{default:p(()=>[s(v,{inset:"",style:{margin:"6px"},title:"\u670D\u52A1\u7AEF\u4FE1\u606F"},{default:p(()=>[s(a,{modelValue:n.value,"onUpdate:modelValue":o[0]||(o[0]=m=>n.value=m),label:"\u670D\u52A1\u5668\u540D",placeholder:"",readonly:""},null,8,["modelValue"]),s(a,{modelValue:t.value,"onUpdate:modelValue":o[1]||(o[1]=m=>t.value=m),label:"\u670D\u52A1\u5668\u7248\u672C",placeholder:"",readonly:""},null,8,["modelValue"]),B("div",k,[s(f,{round:"",block:"",type:"primary",disabled:r.value,"native-type":"submit"},{default:p(()=>[S(" \u66F4\u65B0\u6570\u636E")]),_:1},8,["disabled"])])]),_:1})]),_:1})])}}});const Q=N(G,[["__scopeId","data-v-4ec6fbe9"]]);export{Q as default};
