var A=Object.defineProperty;var F=Object.getOwnPropertySymbols;var C=Object.prototype.hasOwnProperty,B=Object.prototype.propertyIsEnumerable;var f=(n,l,a)=>l in n?A(n,l,{enumerable:!0,configurable:!0,writable:!0,value:a}):n[l]=a,b=(n,l)=>{for(var a in l||(l={}))C.call(l,a)&&f(n,a,l[a]);if(F)for(var a of F(l))B.call(l,a)&&f(n,a,l[a]);return n};var g=(n,l,a)=>new Promise((e,d)=>{var m=u=>{try{i(a.next(u))}catch(s){d(s)}},D=u=>{try{i(a.throw(u))}catch(s){d(s)}},i=u=>u.done?e(u.value):Promise.resolve(u.value).then(m,D);i((a=a.apply(n,l)).next())});import{at as k,au as v,d as S,av as y,r as P,Y as _,O as T,y as U,b as o,C as p,Z as N,a0 as x,D as w,U as t,P as R,H as O}from"./9a74f89d.js";import{C as $}from"./bd857e79.js";import"./74dee749.js";import{F as K}from"./6c55285d.js";/* empty css        */import{_ as q}from"./f155ca5d.js";import{u as G,g as M}from"./4321409c.js";import{F as H}from"./9d312194.js";import{_ as I}from"./0ee26181.js";import"./8f51a641.js";import"./e27e03ce.js";import"./709c511c.js";const V={processName:"",executablePath:"",memoryThreshold:"",checkPeriod:"",restartDelay:"",rconAddress:"",rconPort:"",rconPasswd:"",sourceDir:"",destDir:"",backupTime:"",backupCount:"",accessToken:"",secretKey:"",bucket:""},Y=k({id:"serverSetting",state:()=>({result:"",error:"",formData:b({},V)}),getters:{getResult(){return this.result||""},getError(){return this.error||""},getFormData(){return this.formData||{}}},actions:{updateField(n,l){this.formData&&(this.formData[n]=l)},resetForm(){this.formData=b({},V)},setFormData(n){this.formData=n},SubmitForm(n){return g(this,null,function*(){return this.error="",new Promise((l,a)=>{G(n).then(e=>{const d=e.message||null;(e.code||404)==200?this.result=e.message:this.error=d,l(e)}).catch(e=>{this.error=e,a(e)})})})},GetServerSettingsData(){return g(this,null,function*(){return new Promise((n,l)=>{M().then(a=>{const e=a.message||null;(a.code||404)==200?this.setFormData(a.result||{}):this.error=e,n(a)}).catch(a=>{this.error=a,l(a)})})})}}});function Z(){return Y(v)}const j={class:"page_content"},z={style:{margin:"16px"}},J=S({__name:"ServerSetting",setup(n){const l=Z(),{getFormData:a}=y(l);var e=a;const d=i=>`${i} \u4E0D\u5408\u6CD5\uFF0C\u8BF7\u91CD\u65B0\u8F93\u5165`,m=P(!1),D=i=>g(this,null,function*(){m.value=!0,yield l.SubmitForm(i),l.error?(console.error("Failed to Update:",l.getError),N("\u66F4\u65B0\u670D\u52A1\u5668\u914D\u7F6E\u5931\u8D25: "+l.getError),m.value=!1):(x("\u66F4\u65B0\u670D\u52A1\u5668\u914D\u7F6E\u6210\u529F"),m.value=!1)});return _(()=>{l.GetServerSettingsData()}),(i,u)=>{const s=K,c=$,E=T,h=H;return w(),U("div",j,[o(q),o(h,{onSubmit:D},{default:p(()=>[o(c,{title:"\u670D\u52A1\u7AEF"},{default:p(()=>[o(s,{label:"\u8FDB\u7A0B\u540D",name:"processName","input-align":"right",center:!0,border:!1,modelValue:t(e).processName,"onUpdate:modelValue":u[0]||(u[0]=r=>t(e).processName=r)},null,8,["modelValue"]),o(s,{label:"\u53EF\u6267\u884C\u6587\u4EF6",name:"executablePath","input-align":"right",center:!0,border:!1,modelValue:t(e).executablePath,"onUpdate:modelValue":u[1]||(u[1]=r=>t(e).executablePath=r)},null,8,["modelValue"]),o(s,{label:"\u5185\u5B58\u9608\u503C(%)",name:"memoryThreshold","input-align":"right",center:!0,border:!1,modelValue:t(e).memoryThreshold,"onUpdate:modelValue":u[2]||(u[2]=r=>t(e).memoryThreshold=r),rules:[{required:!0,message:"\u5185\u5B58\u9608\u503C\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:d}]},null,8,["modelValue","rules"]),o(s,{label:"\u68C0\u6D4B\u5468\u671F(\u79D2)",name:"checkPeriod","input-align":"right",center:!0,border:!1,modelValue:t(e).checkPeriod,"onUpdate:modelValue":u[3]||(u[3]=r=>t(e).checkPeriod=r),rules:[{required:!0,message:"\u68C0\u6D4B\u5468\u671F\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:d}]},null,8,["modelValue","rules"]),o(s,{label:"\u91CD\u542F\u5EF6\u8FDF(\u79D2)",name:"restartDelay","input-align":"right",center:!0,border:!1,modelValue:t(e).restartDelay,"onUpdate:modelValue":u[4]||(u[4]=r=>t(e).restartDelay=r),rules:[{required:!0,message:"\u91CD\u542F\u5EF6\u8FDF\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:d}]},null,8,["modelValue","rules"])]),_:1}),o(c,{title:"RCON\u8BBE\u7F6E"},{default:p(()=>[o(s,{label:"RCON\u5730\u5740",name:"rconAddress","input-align":"right",center:!0,border:!1,modelValue:t(e).rconAddress,"onUpdate:modelValue":u[5]||(u[5]=r=>t(e).rconAddress=r),placeholder:"RCON\u5730\u5740"},null,8,["modelValue"]),o(s,{label:"RCON\u7AEF\u53E3",name:"rconPort","input-align":"right",center:!0,border:!1,modelValue:t(e).rconPort,"onUpdate:modelValue":u[6]||(u[6]=r=>t(e).rconPort=r),placeholder:"RCON\u7AEF\u53E3"},null,8,["modelValue"]),o(s,{label:"RCON\u5BC6\u7801",name:"rconPasswd","input-align":"right",center:!0,border:!1,modelValue:t(e).rconPasswd,"onUpdate:modelValue":u[7]||(u[7]=r=>t(e).rconPasswd=r),placeholder:"RCON\u5BC6\u7801"},null,8,["modelValue"])]),_:1}),o(c,{title:"\u672C\u5730\u5907\u4EFD"},{default:p(()=>[o(s,{label:"\u5B58\u6863\u8DEF\u5F84",name:"sourceDir","input-align":"right",center:!0,border:!1,modelValue:t(e).sourceDir,"onUpdate:modelValue":u[8]||(u[8]=r=>t(e).sourceDir=r)},null,8,["modelValue"]),o(s,{label:"\u5907\u4EFD\u8DEF\u5F84",name:"destDir","input-align":"right",center:!0,border:!1,modelValue:t(e).destDir,"onUpdate:modelValue":u[9]||(u[9]=r=>t(e).destDir=r)},null,8,["modelValue"]),o(s,{label:"\u5B58\u6863\u5468\u671F(\u79D2)",name:"backupTime","input-align":"right",center:!0,border:!1,modelValue:t(e).backupTime,"onUpdate:modelValue":u[10]||(u[10]=r=>t(e).backupTime=r),rules:[{required:!0,message:"\u5B58\u6863\u5468\u671F\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:d}]},null,8,["modelValue","rules"]),o(s,{label:"\u6EDA\u52A8\u4FDD\u5B58\u6570",name:"backupCount","input-align":"right",center:!0,border:!1,modelValue:t(e).backupCount,"onUpdate:modelValue":u[11]||(u[11]=r=>t(e).backupCount=r),rules:[{required:!0,message:"\u6EDA\u52A8\u4FDD\u5B58\u6570\u4E0D\u80FD\u4E3A\u7A7A"},{pattern:/^[0-9]+$/,message:d}]},null,8,["modelValue","rules"])]),_:1}),o(c,{title:"\u4E03\u725B\u5907\u4EFD"},{default:p(()=>[o(s,{label:"accessToken",name:"accessToken","input-align":"right",center:!0,border:!1,modelValue:t(e).accessToken,"onUpdate:modelValue":u[12]||(u[12]=r=>t(e).accessToken=r),placeholder:"\u4E03\u725BaccessToken"},null,8,["modelValue"]),o(s,{label:"secretKey",name:"secretKey","input-align":"right",center:!0,border:!1,modelValue:t(e).secretKey,"onUpdate:modelValue":u[13]||(u[13]=r=>t(e).secretKey=r),placeholder:"\u4E03\u725BsecretKey"},null,8,["modelValue"]),o(s,{label:"bucket",name:"bucket","input-align":"right",center:!0,border:!1,modelValue:t(e).bucket,"onUpdate:modelValue":u[14]||(u[14]=r=>t(e).bucket=r),placeholder:"\u4E03\u725Bbucket"},null,8,["modelValue"])]),_:1}),R("div",z,[o(E,{round:"",block:"",type:"primary",disabled:m.value,"native-type":"submit"},{default:p(()=>[O(" \u4FEE\u6539 ")]),_:1},8,["disabled"])])]),_:1})])}}});const de=I(J,[["__scopeId","data-v-85b99b0a"]]);export{de as default};
