import{ax as e}from"./9da9b27d.js";function t(r){return e.request({url:"/BroadCast/send",method:"POST",params:r},{isTransformResponse:!1})}function n(){return e.request({url:"/Players/show",method:"GET"},{isTransformResponse:!1})}function o(){return e.request({url:"/rcon/info",method:"GET"},{isTransformResponse:!1})}function a(r){return e.request({url:"/rcon/restart",method:"GET",params:r},{isTransformResponse:!1})}function u(r){return e.request({url:"/serverSetting/update",method:"POST",params:r},{isTransformResponse:!1})}function f(){return e.request({url:"/serverSetting/get",method:"GET"},{isTransformResponse:!1})}export{n as a,f as g,o as i,a as r,t as s,u};
