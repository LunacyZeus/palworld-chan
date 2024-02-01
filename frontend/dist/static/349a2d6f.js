import{c as B,t as p,n as x,m as S,d as _,u as y,b as t,W as b,az as I,w as P,f as R,h as w,k as N,i as T,l as v,aA as $,p as q,o as z,I as O,v as G,A as V,D as M,a5 as U}from"./9da9b27d.js";const[h,j]=B("grid"),K={square:Boolean,center:p,border:p,gutter:x,reverse:Boolean,iconSize:x,direction:String,clickable:Boolean,columnNum:S(4)},A=Symbol(h);var L=_({name:h,props:K,setup(e,{slots:n}){const{linkChildren:r}=y(A);return r({props:e}),()=>{var c;return t("div",{style:{paddingLeft:b(e.gutter)},class:[j(),{[I]:e.border&&!e.gutter}]},[(c=n.default)==null?void 0:c.call(n)])}}});const W=P(L),[Y,g]=B("grid-item"),H=R({},w,{dot:Boolean,text:String,icon:String,badge:x,iconColor:String,iconPrefix:String,badgeProps:Object});var J=_({name:Y,props:H,setup(e,{slots:n}){const{parent:r,index:c}=N(A),o=T();if(!r)return;const m=v(()=>{const{square:s,gutter:a,columnNum:u}=r.props,d=`${100/+u}%`,i={flexBasis:d};if(s)i.paddingTop=d;else if(a){const l=b(a);i.paddingRight=l,c.value>=u&&(i.marginTop=l)}return i}),D=v(()=>{const{square:s,gutter:a}=r.props;if(s&&a){const u=b(a);return{right:u,bottom:u,height:"auto"}}}),E=()=>{if(n.icon)return t(z,q({dot:e.dot,content:e.badge},e.badgeProps),{default:n.icon});if(e.icon)return t(O,{dot:e.dot,name:e.icon,size:r.props.iconSize,badge:e.badge,class:g("icon"),color:e.iconColor,badgeProps:e.badgeProps,classPrefix:e.iconPrefix},null)},C=()=>{if(n.text)return n.text();if(e.text)return t("span",{class:g("text")},[e.text])},F=()=>n.default?n.default():[E(),C()];return()=>{const{center:s,border:a,square:u,gutter:d,reverse:i,direction:l,clickable:f}=r.props,k=[g("content",[l,{center:s,square:u,reverse:i,clickable:f,surround:a&&d}]),{[$]:a}];return t("div",{class:[g({square:u})],style:m.value},[t("div",{role:f?"button":void 0,class:k,style:D.value,tabindex:f?0:void 0,onClick:o},[F()])])}}});const Q=P(J);const X={id:"app"},ee=_({__name:"index",setup(e){const n=()=>{U("\u672A\u5B9E\u88C5")};return(r,c)=>{const o=Q,m=W;return M(),G("div",X,[t(m,{clickable:"","column-num":3},{default:V(()=>[t(o,{icon:"bullhorn-o",text:"\u670D\u52A1\u5668\u5E7F\u64AD",to:"/broadCast"}),t(o,{icon:"friends-o",text:"\u5728\u7EBF\u7528\u6237",to:"/online"}),t(o,{icon:"bullhorn-o",text:"\u5B58\u6863\u7BA1\u7406",to:"/saveManager"}),t(o,{icon:"records",text:"\u5907\u4EFD\u7BA1\u7406",to:"/backUpManager"}),t(o,{icon:"notes-o",text:"\u670D\u52A1\u7AEF\u914D\u7F6E",onClick:n}),t(o,{icon:"info-o",text:"\u670D\u52A1\u7AEF\u4FE1\u606F",to:"/gameServer/info"}),t(o,{icon:"replay",text:"\u91CD\u542F\u670D\u52A1\u7AEF",to:"/gameServer/restart"}),t(o,{icon:"underway-o",text:"\u5B9A\u65F6\u5E7F\u64AD",onClick:n})]),_:1})])}}});export{ee as default};
