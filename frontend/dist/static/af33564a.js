import{ad as le,f as G,aE as A,c as J,d as X,n as B,V as k,_ as F,r as D,aF as oe,k as ie,Y as Q,aG as ue,b as m,ao as W,m as _,t as ce,U as se,u as re,l as x,aH as de,X as p,aI as R,ac as me,M as K,aJ as ve,w as fe}from"./373881a5.js";const Z=e=>e.find(l=>!l.disabled)||e[0];function he(e,l){const n=e[0];if(n){if(Array.isArray(n))return"multiple";if(l.children in n)return"cascade"}return"default"}function U(e,l){l=A(l,0,e.length);for(let n=l;n<e.length;n++)if(!e[n].disabled)return n;for(let n=l-1;n>=0;n--)if(!e[n].disabled)return n;return 0}const L=(e,l,n)=>l!==void 0&&!!e.find(o=>o[n.value]===l);function Y(e,l,n){const o=e.findIndex(v=>v[n.value]===l),r=U(e,o);return e[r]}function ge(e,l,n){const o=[];let r={[l.children]:e},v=0;for(;r&&r[l.children];){const c=r[l.children],d=n.value[v];if(r=le(d)?Y(c,d,l):void 0,!r&&c.length){const T=Z(c)[l.value];r=Y(c,T,l)}v++,o.push(c)}return o}function be(e){const{transform:l}=window.getComputedStyle(e),n=l.slice(7,l.length-1).split(", ")[5];return Number(n)}function Te(e){return G({text:"text",value:"value",children:"children"},e)}const j=200,q=300,Oe=15,[ee,$]=J("picker-column"),te=Symbol(ee);var ye=X({name:ee,props:{value:B,fields:k(Object),options:F(),readonly:Boolean,allowHtml:Boolean,optionHeight:k(Number),swipeDuration:k(B),visibleOptionNum:k(B)},emits:["change"],setup(e,{emit:l,slots:n}){let o,r,v,c,d;const T=D(),s=D(0),f=D(0),h=oe(),O=()=>e.options.length,C=()=>e.optionHeight*(+e.visibleOptionNum-1)/2,y=i=>{const t=U(e.options,i),a=-t*e.optionHeight,u=()=>{const b=e.options[t][e.fields.value];b!==e.value&&l("change",b)};o&&a!==s.value?d=u:u(),s.value=a},N=i=>{o||e.readonly||(d=null,f.value=j,y(i))},E=i=>A(Math.round(-i/e.optionHeight),0,O()-1),P=(i,t)=>{const a=Math.abs(i/t);i=s.value+a/.003*(i<0?-1:1);const u=E(i);f.value=+e.swipeDuration,y(u)},H=()=>{o=!1,f.value=0,d&&(d(),d=null)},M=i=>{if(!e.readonly){if(h.start(i),o){const t=be(T.value);s.value=Math.min(0,t-C())}f.value=0,r=s.value,v=Date.now(),c=r,d=null}},I=i=>{if(e.readonly)return;h.move(i),h.isVertical()&&(o=!0,W(i,!0)),s.value=A(r+h.deltaY.value,-(O()*e.optionHeight),e.optionHeight);const t=Date.now();t-v>q&&(v=t,c=s.value)},S=()=>{if(e.readonly)return;const i=s.value-c,t=Date.now()-v;if(t<q&&Math.abs(i)>Oe){P(i,t);return}const u=E(s.value);f.value=j,y(u),setTimeout(()=>{o=!1},0)},V=()=>{const i={height:`${e.optionHeight}px`};return e.options.map((t,a)=>{const u=t[e.fields.text],{disabled:b}=t,w=t[e.fields.value],ne={role:"button",style:i,tabindex:b?-1:0,class:[$("item",{disabled:b,selected:w===e.value}),t.className],onClick:()=>N(a)},ae={class:"van-ellipsis",[e.allowHtml?"innerHTML":"textContent"]:u};return m("li",ne,[n.option?n.option(t):m("div",ae,null)])})};return ie(te),Q({stopMomentum:H}),ue(()=>{const i=e.options.findIndex(u=>u[e.fields.value]===e.value),a=-U(e.options,i)*e.optionHeight;s.value=a}),()=>m("div",{class:$(),onTouchstart:M,onTouchmove:I,onTouchend:S,onTouchcancel:S},[m("ul",{ref:T,style:{transform:`translate3d(0, ${s.value+C()}px, 0)`,transitionDuration:`${f.value}ms`,transitionProperty:f.value?"all":"none"},class:$("wrapper"),onTransitionend:H},[V()])])}});const[xe,g,z]=J("picker"),we={title:String,loading:Boolean,readonly:Boolean,allowHtml:Boolean,optionHeight:_(44),showToolbar:ce,swipeDuration:_(1e3),visibleOptionNum:_(6),cancelButtonText:String,confirmButtonText:String},Ce=G({},we,{columns:F(),modelValue:F(),toolbarPosition:se("top"),columnsFieldNames:Object});var Ee=X({name:xe,props:Ce,emits:["confirm","cancel","change","update:modelValue"],setup(e,{emit:l,slots:n}){const o=D(e.modelValue),{children:r,linkChildren:v}=re(te);v();const c=x(()=>Te(e.columnsFieldNames)),d=x(()=>de(e.optionHeight)),T=x(()=>he(e.columns,c.value)),s=x(()=>{const{columns:t}=e;switch(T.value){case"multiple":return t;case"cascade":return ge(t,c.value,o);default:return[t]}}),f=x(()=>s.value.some(t=>t.length)),h=x(()=>s.value.map((t,a)=>Y(t,o.value[a],c.value))),O=(t,a)=>{if(o.value[t]!==a){const u=o.value.slice(0);u[t]=a,o.value=u}},C=(t,a)=>{O(a,t),T.value==="cascade"&&o.value.forEach((u,b)=>{const w=s.value[b];L(w,u,c.value)||O(b,w.length?w[0][c.value.value]:void 0)}),l("change",{columnIndex:a,selectedValues:o.value,selectedOptions:h.value})},y=()=>{r.forEach(t=>t.stopMomentum()),l("confirm",{selectedValues:o.value,selectedOptions:h.value})},N=()=>l("cancel",{selectedValues:o.value,selectedOptions:h.value}),E=()=>{if(n.title)return n.title();if(e.title)return m("div",{class:[g("title"),"van-ellipsis"]},[e.title])},P=()=>{const t=e.cancelButtonText||z("cancel");return m("button",{type:"button",class:[g("cancel"),K],onClick:N},[n.cancel?n.cancel():t])},H=()=>{const t=e.confirmButtonText||z("confirm");return m("button",{type:"button",class:[g("confirm"),K],onClick:y},[n.confirm?n.confirm():t])},M=()=>{if(e.showToolbar)return m("div",{class:g("toolbar")},[n.toolbar?n.toolbar():[P(),E(),H()]])},I=()=>s.value.map((t,a)=>m(ye,{value:o.value[a],fields:c.value,options:t,readonly:e.readonly,allowHtml:e.allowHtml,optionHeight:d.value,swipeDuration:e.swipeDuration,visibleOptionNum:e.visibleOptionNum,onChange:u=>C(u,a)},{option:n.option})),S=t=>{if(f.value){const a={height:`${d.value}px`},u={backgroundSize:`100% ${(t-d.value)/2}px`};return[m("div",{class:g("mask"),style:u},null),m("div",{class:[ve,g("frame")],style:a},null)]}},V=()=>{const t=d.value*+e.visibleOptionNum,a={height:`${t}px`};return m("div",{class:g("columns"),style:a,onTouchmove:W},[I(),S(t)])};return p(s,t=>{t.forEach((a,u)=>{a.length&&!L(a,o.value[u],c.value)&&O(u,Z(a)[c.value.value])})},{immediate:!0}),p(()=>e.modelValue,t=>{R(t,o.value)||(o.value=t)},{deep:!0}),p(o,t=>{R(t,e.modelValue)||l("update:modelValue",t)},{immediate:!0}),Q({confirm:y,getSelectedOptions:()=>h.value}),()=>{var t,a;return m("div",{class:g()},[e.toolbarPosition==="top"?M():null,e.loading?m(me,{class:g("loading")},null):null,(t=n["columns-top"])==null?void 0:t.call(n),V(),(a=n["columns-bottom"])==null?void 0:a.call(n),e.toolbarPosition==="bottom"?M():null])}}});const Me=fe(Ee);export{Me as P};
