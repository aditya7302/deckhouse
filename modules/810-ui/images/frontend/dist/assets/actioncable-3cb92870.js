import{c as M,a as _e,b as Te,l as Ae}from"./index-200c7ba5.js";import{d as Y}from"./deckhouse_settings-2c9d1e33.js";function Me(v){throw new Error('Could not dynamically require "'+v+'". Please configure the dynamicRequireTargets or/and ignoreDynamicRequires option of @rollup/plugin-commonjs appropriately for this require call to work.')}var I={},We={get exports(){return I},set exports(v){I=v}};(function(v,L){(function(R,y){y(L)})(M,function(R){var y=typeof window<"u"?window:typeof M<"u"?M:typeof self<"u"?self:{},_=function(t,n){if(n=n.split(":")[0],t=+t,!t)return!1;switch(n){case"http":case"ws":return t!==80;case"https":case"wss":return t!==443;case"ftp":return t!==21;case"gopher":return t!==70;case"file":return!1}return t!==0},G=Object.prototype.hasOwnProperty,ie;function Z(r){try{return decodeURIComponent(r.replace(/\+/g," "))}catch{return null}}function J(r){try{return encodeURIComponent(r)}catch{return null}}function se(r){for(var t=/([^=?#&]+)=?([^&]*)/g,n={},e;e=t.exec(r);){var o=Z(e[1]),i=Z(e[2]);o===null||i===null||o in n||(n[o]=i)}return n}function ae(r,t){t=t||"";var n=[],e,o;typeof t!="string"&&(t="?");for(o in r)if(G.call(r,o)){if(e=r[o],!e&&(e===null||e===ie||isNaN(e))&&(e=""),o=J(o),e=J(e),o===null||e===null)continue;n.push(o+"="+e)}return n.length?t+n.join("&"):""}var ce=ae,ue=se,U={stringify:ce,parse:ue},X=/[\n\r\t]/g,le=/^[A-Za-z][A-Za-z0-9+-.]*:\/\//,he=/^([a-z][a-z0-9.+-]*:)?(\/\/)?([\\/]+)?([\S\s]*)/i,fe=/^[a-zA-Z]:/,pe=/^[\x00-\x20\u00a0\u1680\u2000-\u200a\u2028\u2029\u202f\u205f\u3000\ufeff]+/;function P(r){return(r||"").toString().replace(pe,"")}var q=[["#","hash"],["?","query"],function(t,n){return b(n.protocol)?t.replace(/\\/g,"/"):t},["/","pathname"],["@","auth",1],[NaN,"host",void 0,1,1],[/:(\d*)$/,"port",void 0,1],[NaN,"hostname",void 0,1,1]],K={hash:1,query:1};function Q(r){var t;typeof window<"u"?t=window:typeof y<"u"?t=y:typeof self<"u"?t=self:t={};var n=t.location||{};r=r||n;var e={},o=typeof r,i;if(r.protocol==="blob:")e=new O(unescape(r.pathname),{});else if(o==="string"){e=new O(r,{});for(i in K)delete e[i]}else if(o==="object"){for(i in r)i in K||(e[i]=r[i]);e.slashes===void 0&&(e.slashes=le.test(r.href))}return e}function b(r){return r==="file:"||r==="ftp:"||r==="http:"||r==="https:"||r==="ws:"||r==="wss:"}function $(r,t){r=P(r),r=r.replace(X,""),t=t||{};var n=he.exec(r),e=n[1]?n[1].toLowerCase():"",o=!!n[2],i=!!n[3],s=0,a;return o?i?(a=n[2]+n[3]+n[4],s=n[2].length+n[3].length):(a=n[2]+n[4],s=n[2].length):i?(a=n[3]+n[4],s=n[3].length):a=n[4],e==="file:"?s>=2&&(a=a.slice(2)):b(e)?a=n[4]:e?o&&(a=a.slice(2)):s>=2&&b(t.protocol)&&(a=n[4]),{protocol:e,slashes:o||b(e),slashesCount:s,rest:a}}function ve(r,t){if(r==="")return t;for(var n=(t||"/").split("/").slice(0,-1).concat(r.split("/")),e=n.length,o=n[e-1],i=!1,s=0;e--;)n[e]==="."?n.splice(e,1):n[e]===".."?(n.splice(e,1),s++):s&&(e===0&&(i=!0),n.splice(e,1),s--);return i&&n.unshift(""),(o==="."||o==="..")&&n.push(""),n.join("/")}function O(r,t,n){if(r=P(r),r=r.replace(X,""),!(this instanceof O))return new O(r,t,n);var e,o,i,s,a,u,l=q.slice(),p=typeof t,c=this,H=0;for(p!=="object"&&p!=="string"&&(n=t,t=null),n&&typeof n!="function"&&(n=U.parse),t=Q(t),o=$(r||"",t),e=!o.protocol&&!o.slashes,c.slashes=o.slashes||e&&t.slashes,c.protocol=o.protocol||t.protocol||"",r=o.rest,(o.protocol==="file:"&&(o.slashesCount!==2||fe.test(r))||!o.slashes&&(o.protocol||o.slashesCount<2||!b(c.protocol)))&&(l[3]=[/(.*)/,"pathname"]);H<l.length;H++){if(s=l[H],typeof s=="function"){r=s(r,c);continue}i=s[0],u=s[1],i!==i?c[u]=r:typeof i=="string"?(a=i==="@"?r.lastIndexOf(i):r.indexOf(i),~a&&(typeof s[2]=="number"?(c[u]=r.slice(0,a),r=r.slice(a+s[2])):(c[u]=r.slice(a),r=r.slice(0,a)))):(a=i.exec(r))&&(c[u]=a[1],r=r.slice(0,a.index)),c[u]=c[u]||e&&s[3]&&t[u]||"",s[4]&&(c[u]=c[u].toLowerCase())}n&&(c.query=n(c.query)),e&&t.slashes&&c.pathname.charAt(0)!=="/"&&(c.pathname!==""||t.pathname!=="")&&(c.pathname=ve(c.pathname,t.pathname)),c.pathname.charAt(0)!=="/"&&b(c.protocol)&&(c.pathname="/"+c.pathname),_(c.port,c.protocol)||(c.host=c.hostname,c.port=""),c.username=c.password="",c.auth&&(a=c.auth.indexOf(":"),~a?(c.username=c.auth.slice(0,a),c.username=encodeURIComponent(decodeURIComponent(c.username)),c.password=c.auth.slice(a+1),c.password=encodeURIComponent(decodeURIComponent(c.password))):c.username=encodeURIComponent(decodeURIComponent(c.auth)),c.auth=c.password?c.username+":"+c.password:c.username),c.origin=c.protocol!=="file:"&&b(c.protocol)&&c.host?c.protocol+"//"+c.host:"null",c.href=c.toString()}function de(r,t,n){var e=this;switch(r){case"query":typeof t=="string"&&t.length&&(t=(n||U.parse)(t)),e[r]=t;break;case"port":e[r]=t,_(t,e.protocol)?t&&(e.host=e.hostname+":"+t):(e.host=e.hostname,e[r]="");break;case"hostname":e[r]=t,e.port&&(t+=":"+e.port),e.host=t;break;case"host":e[r]=t,/:\d+$/.test(t)?(t=t.split(":"),e.port=t.pop(),e.hostname=t.join(":")):(e.hostname=t,e.port="");break;case"protocol":e.protocol=t.toLowerCase(),e.slashes=!n;break;case"pathname":case"hash":if(t){var o=r==="pathname"?"/":"#";e[r]=t.charAt(0)!==o?o+t:t}else e[r]=t;break;case"username":case"password":e[r]=encodeURIComponent(t);break;case"auth":var i=t.indexOf(":");~i?(e.username=t.slice(0,i),e.username=encodeURIComponent(decodeURIComponent(e.username)),e.password=t.slice(i+1),e.password=encodeURIComponent(decodeURIComponent(e.password))):e.username=encodeURIComponent(decodeURIComponent(t))}for(var s=0;s<q.length;s++){var a=q[s];a[4]&&(e[a[1]]=e[a[1]].toLowerCase())}return e.auth=e.password?e.username+":"+e.password:e.username,e.origin=e.protocol!=="file:"&&b(e.protocol)&&e.host?e.protocol+"//"+e.host:"null",e.href=e.toString(),e}function ye(r){(!r||typeof r!="function")&&(r=U.stringify);var t,n=this,e=n.host,o=n.protocol;o&&o.charAt(o.length-1)!==":"&&(o+=":");var i=o+(n.protocol&&n.slashes||b(n.protocol)?"//":"");return n.username?(i+=n.username,n.password&&(i+=":"+n.password),i+="@"):n.password?(i+=":"+n.password,i+="@"):n.protocol!=="file:"&&b(n.protocol)&&!e&&n.pathname!=="/"&&(i+="@"),e[e.length-1]===":"&&(e+=":"),i+=e+n.pathname,t=typeof n.query=="object"?r(n.query):n.query,t&&(i+=t.charAt(0)!=="?"?"?"+t:t),n.hash&&(i+=n.hash),i}O.prototype={set:de,toString:ye},O.extractProtocol=$,O.location=Q,O.trimLeft=P,O.qs=U;var B=O;function T(r,t){setTimeout(function(n){return r.call(n)},4,t)}function j(r,t){typeof process<"u"&&console[r].call(null,t)}function F(r,t){r===void 0&&(r=[]);var n=[];return r.forEach(function(e){t(e)||n.push(e)}),n}function me(r,t){r===void 0&&(r=[]);var n=[];return r.forEach(function(e){t(e)&&n.push(e)}),n}var C=function(){this.listeners={}};C.prototype.addEventListener=function(t,n){typeof n=="function"&&(Array.isArray(this.listeners[t])||(this.listeners[t]=[]),me(this.listeners[t],function(e){return e===n}).length===0&&this.listeners[t].push(n))},C.prototype.removeEventListener=function(t,n){var e=this.listeners[t];this.listeners[t]=F(e,function(o){return o===n})},C.prototype.dispatchEvent=function(t){for(var n=this,e=[],o=arguments.length-1;o-- >0;)e[o]=arguments[o+1];var i=t.type,s=this.listeners[i];return Array.isArray(s)?(s.forEach(function(a){e.length>0?a.apply(n,e):a.call(n,t)}),!0):!1};function w(r){var t=r.indexOf("?");return t>=0?r.slice(0,t):r}var g=function(){this.urlMap={}};g.prototype.attachWebSocket=function(t,n){var e=w(n),o=this.urlMap[e];if(o&&o.server&&o.websockets.indexOf(t)===-1)return o.websockets.push(t),o.server},g.prototype.addMembershipToRoom=function(t,n){var e=this.urlMap[w(t.url)];e&&e.server&&e.websockets.indexOf(t)!==-1&&(e.roomMemberships[n]||(e.roomMemberships[n]=[]),e.roomMemberships[n].push(t))},g.prototype.attachServer=function(t,n){var e=w(n),o=this.urlMap[e];if(!o)return this.urlMap[e]={server:t,websockets:[],roomMemberships:{}},t},g.prototype.serverLookup=function(t){var n=w(t),e=this.urlMap[n];if(e)return e.server},g.prototype.websocketsLookup=function(t,n,e){var o=w(t),i,s=this.urlMap[o];if(i=s?s.websockets:[],n){var a=s.roomMemberships[n];i=a||[]}return e?i.filter(function(u){return u!==e}):i},g.prototype.removeServer=function(t){delete this.urlMap[w(t)]},g.prototype.removeWebSocket=function(t,n){var e=w(n),o=this.urlMap[e];o&&(o.websockets=F(o.websockets,function(i){return i===t}))},g.prototype.removeMembershipFromRoom=function(t,n){var e=this.urlMap[w(t.url)],o=e.roomMemberships[n];e&&o!==null&&(e.roomMemberships[n]=F(o,function(i){return i===t}))};var h=new g,m={CLOSE_NORMAL:1e3,CLOSE_GOING_AWAY:1001,CLOSE_PROTOCOL_ERROR:1002,CLOSE_UNSUPPORTED:1003,CLOSE_NO_STATUS:1005,CLOSE_ABNORMAL:1006,UNSUPPORTED_DATA:1007,POLICY_VIOLATION:1008,CLOSE_TOO_LARGE:1009,MISSING_EXTENSION:1010,INTERNAL_ERROR:1011,SERVICE_RESTART:1012,TRY_AGAIN_LATER:1013,TLS_HANDSHAKE:1015},d={CONSTRUCTOR_ERROR:"Failed to construct 'WebSocket':",CLOSE_ERROR:"Failed to execute 'close' on 'WebSocket':",EVENT:{CONSTRUCT:"Failed to construct 'Event':",MESSAGE:"Failed to construct 'MessageEvent':",CLOSE:"Failed to construct 'CloseEvent':"}},k=function(){};k.prototype.stopPropagation=function(){},k.prototype.stopImmediatePropagation=function(){},k.prototype.initEvent=function(t,n,e){t===void 0&&(t="undefined"),n===void 0&&(n=!1),e===void 0&&(e=!1),this.type=""+t,this.bubbles=Boolean(n),this.cancelable=Boolean(e)};var Ee=function(r){function t(n,e){if(e===void 0&&(e={}),r.call(this),!n)throw new TypeError(d.EVENT_ERROR+" 1 argument required, but only 0 present.");if(typeof e!="object")throw new TypeError(d.EVENT_ERROR+" parameter 2 ('eventInitDict') is not an object.");var o=e.bubbles,i=e.cancelable;this.type=""+n,this.timeStamp=Date.now(),this.target=null,this.srcElement=null,this.returnValue=!0,this.isTrusted=!1,this.eventPhase=0,this.defaultPrevented=!1,this.currentTarget=null,this.cancelable=i?Boolean(i):!1,this.cancelBubble=!1,this.bubbles=o?Boolean(o):!1}return r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t,t}(k),Se=function(r){function t(n,e){if(e===void 0&&(e={}),r.call(this),!n)throw new TypeError(d.EVENT.MESSAGE+" 1 argument required, but only 0 present.");if(typeof e!="object")throw new TypeError(d.EVENT.MESSAGE+" parameter 2 ('eventInitDict') is not an object");var o=e.bubbles,i=e.cancelable,s=e.data,a=e.origin,u=e.lastEventId,l=e.ports;this.type=""+n,this.timeStamp=Date.now(),this.target=null,this.srcElement=null,this.returnValue=!0,this.isTrusted=!1,this.eventPhase=0,this.defaultPrevented=!1,this.currentTarget=null,this.cancelable=i?Boolean(i):!1,this.canncelBubble=!1,this.bubbles=o?Boolean(o):!1,this.origin=""+a,this.ports=typeof l>"u"?null:l,this.data=typeof s>"u"?null:s,this.lastEventId=""+(u||"")}return r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t,t}(k),be=function(r){function t(n,e){if(e===void 0&&(e={}),r.call(this),!n)throw new TypeError(d.EVENT.CLOSE+" 1 argument required, but only 0 present.");if(typeof e!="object")throw new TypeError(d.EVENT.CLOSE+" parameter 2 ('eventInitDict') is not an object");var o=e.bubbles,i=e.cancelable,s=e.code,a=e.reason,u=e.wasClean;this.type=""+n,this.timeStamp=Date.now(),this.target=null,this.srcElement=null,this.returnValue=!0,this.isTrusted=!1,this.eventPhase=0,this.defaultPrevented=!1,this.currentTarget=null,this.cancelable=i?Boolean(i):!1,this.cancelBubble=!1,this.bubbles=o?Boolean(o):!1,this.code=typeof s=="number"?parseInt(s,10):0,this.reason=""+(a||""),this.wasClean=u?Boolean(u):!1}return r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t,t}(k);function E(r){var t=r.type,n=r.target,e=new Ee(t);return n&&(e.target=n,e.srcElement=n,e.currentTarget=n),e}function A(r){var t=r.type,n=r.origin,e=r.data,o=r.target,i=new Se(t,{data:e,origin:n});return o&&(i.target=o,i.srcElement=o,i.currentTarget=o),i}function S(r){var t=r.code,n=r.reason,e=r.type,o=r.target,i=r.wasClean;i||(i=t===m.CLOSE_NORMAL||t===m.CLOSE_NO_STATUS);var s=new be(e,{code:t,reason:n,wasClean:i});return o&&(s.target=o,s.srcElement=o,s.currentTarget=o),s}function ee(r,t,n){r.readyState=f.CLOSING;var e=h.serverLookup(r.url),o=S({type:"close",target:r.target,code:t,reason:n});T(function(){h.removeWebSocket(r,r.url),r.readyState=f.CLOSED,r.dispatchEvent(o),e&&e.dispatchEvent(o,e)},r)}function Oe(r,t,n){r.readyState=f.CLOSING;var e=h.serverLookup(r.url),o=S({type:"close",target:r.target,code:t,reason:n,wasClean:!1}),i=E({type:"error",target:r.target});T(function(){h.removeWebSocket(r,r.url),r.readyState=f.CLOSED,r.dispatchEvent(i),r.dispatchEvent(o),e&&e.dispatchEvent(o,e)},r)}function D(r){return Object.prototype.toString.call(r)!=="[object Blob]"&&!(r instanceof ArrayBuffer)&&(r=String(r)),r}var V=new WeakMap;function te(r){if(V.has(r))return V.get(r);var t=new Proxy(r,{get:function(e,o){if(o==="close")return function(a){a===void 0&&(a={});var u=a.code||m.CLOSE_NORMAL,l=a.reason||"";ee(t,u,l)};if(o==="send")return function(a){a=D(a),r.dispatchEvent(A({type:"message",data:a,origin:this.url,target:r}))};var i=function(s){return s==="message"?"server::"+s:s};return o==="on"?function(a,u){r.addEventListener(i(a),u)}:o==="off"?function(a,u){r.removeEventListener(i(a),u)}:o==="target"?r:e[o]}});return V.set(r,t),t}function ge(r){var t=encodeURIComponent(r).match(/%[89ABab]/g);return r.length+(t?t.length:0)}function we(r){var t=new B(r),n=t.pathname,e=t.protocol,o=t.hash;if(!r)throw new TypeError(d.CONSTRUCTOR_ERROR+" 1 argument required, but only 0 present.");if(n||(t.pathname="/"),e==="")throw new SyntaxError(d.CONSTRUCTOR_ERROR+" The URL '"+t.toString()+"' is invalid.");if(e!=="ws:"&&e!=="wss:")throw new SyntaxError(d.CONSTRUCTOR_ERROR+" The URL's scheme must be either 'ws' or 'wss'. '"+e+"' is not allowed.");if(o!=="")throw new SyntaxError(d.CONSTRUCTOR_ERROR+" The URL contains a fragment identifier ('"+o+"'). Fragment identifiers are not allowed in WebSocket URLs.");return t.toString()}function Le(r){if(r===void 0&&(r=[]),!Array.isArray(r)&&typeof r!="string")throw new SyntaxError(d.CONSTRUCTOR_ERROR+" The subprotocol '"+r.toString()+"' is invalid.");typeof r=="string"&&(r=[r]);var t=r.map(function(e){return{count:1,protocol:e}}).reduce(function(e,o){return e[o.protocol]=(e[o.protocol]||0)+o.count,e},{}),n=Object.keys(t).filter(function(e){return t[e]>1});if(n.length>0)throw new SyntaxError(d.CONSTRUCTOR_ERROR+" The subprotocol '"+n[0]+"' is duplicated.");return r}var f=function(r){function t(e,o){r.call(this),this._onopen=null,this._onmessage=null,this._onerror=null,this._onclose=null,this.url=we(e),o=Le(o),this.protocol=o[0]||"",this.binaryType="blob",this.readyState=t.CONNECTING;var i=te(this),s=h.attachWebSocket(i,this.url);T(function(){if(s)if(s.options.verifyClient&&typeof s.options.verifyClient=="function"&&!s.options.verifyClient())this.readyState=t.CLOSED,j("error","WebSocket connection to '"+this.url+"' failed: HTTP Authentication failed; no valid credentials available"),h.removeWebSocket(i,this.url),this.dispatchEvent(E({type:"error",target:this})),this.dispatchEvent(S({type:"close",target:this,code:m.CLOSE_NORMAL}));else{if(s.options.selectProtocol&&typeof s.options.selectProtocol=="function"){var u=s.options.selectProtocol(o),l=u!=="",p=o.indexOf(u)!==-1;if(l&&!p){this.readyState=t.CLOSED,j("error","WebSocket connection to '"+this.url+"' failed: Invalid Sub-Protocol"),h.removeWebSocket(i,this.url),this.dispatchEvent(E({type:"error",target:this})),this.dispatchEvent(S({type:"close",target:this,code:m.CLOSE_NORMAL}));return}this.protocol=u}this.readyState=t.OPEN,this.dispatchEvent(E({type:"open",target:this})),s.dispatchEvent(E({type:"connection"}),i)}else this.readyState=t.CLOSED,this.dispatchEvent(E({type:"error",target:this})),this.dispatchEvent(S({type:"close",target:this,code:m.CLOSE_NORMAL})),j("error","WebSocket connection to '"+this.url+"' failed")},this)}r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t;var n={onopen:{},onmessage:{},onclose:{},onerror:{}};return n.onopen.get=function(){return this._onopen},n.onmessage.get=function(){return this._onmessage},n.onclose.get=function(){return this._onclose},n.onerror.get=function(){return this._onerror},n.onopen.set=function(e){this.removeEventListener("open",this._onopen),this._onopen=e,this.addEventListener("open",e)},n.onmessage.set=function(e){this.removeEventListener("message",this._onmessage),this._onmessage=e,this.addEventListener("message",e)},n.onclose.set=function(e){this.removeEventListener("close",this._onclose),this._onclose=e,this.addEventListener("close",e)},n.onerror.set=function(e){this.removeEventListener("error",this._onerror),this._onerror=e,this.addEventListener("error",e)},t.prototype.send=function(o){var i=this;if(this.readyState===t.CLOSING||this.readyState===t.CLOSED)throw new Error("WebSocket is already in CLOSING or CLOSED state");var s=A({type:"server::message",origin:this.url,data:D(o)}),a=h.serverLookup(this.url);a&&T(function(){i.dispatchEvent(s,o)},a)},t.prototype.close=function(o,i){if(o!==void 0&&(typeof o!="number"||o!==1e3&&(o<3e3||o>4999)))throw new TypeError(d.CLOSE_ERROR+" The code must be either 1000, or between 3000 and 4999. "+o+" is neither.");if(i!==void 0){var s=ge(i);if(s>123)throw new SyntaxError(d.CLOSE_ERROR+" The message must not be greater than 123 bytes.")}if(!(this.readyState===t.CLOSING||this.readyState===t.CLOSED)){var a=te(this);this.readyState===t.CONNECTING?Oe(a,o||m.CLOSE_ABNORMAL,i):ee(a,o||m.CLOSE_NO_STATUS,i)}},Object.defineProperties(t.prototype,n),t}(C);f.CONNECTING=0,f.prototype.CONNECTING=f.CONNECTING,f.OPEN=1,f.prototype.OPEN=f.OPEN,f.CLOSING=2,f.prototype.CLOSING=f.CLOSING,f.CLOSED=3,f.prototype.CLOSED=f.CLOSED;var N=function(r){function t(e,o){var i=this;e===void 0&&(e="socket.io"),o===void 0&&(o=""),r.call(this),this.binaryType="blob";var s=new B(e);s.pathname||(s.pathname="/"),this.url=s.toString(),this.readyState=t.CONNECTING,this.protocol="",this.target=this,typeof o=="string"||typeof o=="object"&&o!==null?this.protocol=o:Array.isArray(o)&&o.length>0&&(this.protocol=o[0]);var a=h.attachWebSocket(this,this.url);T(function(){a?(this.readyState=t.OPEN,a.dispatchEvent(E({type:"connection"}),a,this),a.dispatchEvent(E({type:"connect"}),a,this),this.dispatchEvent(E({type:"connect",target:this}))):(this.readyState=t.CLOSED,this.dispatchEvent(E({type:"error",target:this})),this.dispatchEvent(S({type:"close",target:this,code:m.CLOSE_NORMAL})),j("error","Socket.io connection to '"+this.url+"' failed"))},this),this.addEventListener("close",function(u){i.dispatchEvent(S({type:"disconnect",target:u.target,code:u.code}))})}r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t;var n={broadcast:{}};return t.prototype.close=function(){if(this.readyState===t.OPEN){var o=h.serverLookup(this.url);return h.removeWebSocket(this,this.url),this.readyState=t.CLOSED,this.dispatchEvent(S({type:"close",target:this,code:m.CLOSE_NORMAL})),o&&o.dispatchEvent(S({type:"disconnect",target:this,code:m.CLOSE_NORMAL}),o),this}},t.prototype.disconnect=function(){return this.close()},t.prototype.emit=function(o){for(var i=[],s=arguments.length-1;s-- >0;)i[s]=arguments[s+1];if(this.readyState!==t.OPEN)throw new Error("SocketIO is already in CLOSING or CLOSED state");var a=A({type:o,origin:this.url,data:i}),u=h.serverLookup(this.url);return u&&u.dispatchEvent.apply(u,[a].concat(i)),this},t.prototype.send=function(o){return this.emit("message",o),this},n.broadcast.get=function(){if(this.readyState!==t.OPEN)throw new Error("SocketIO is already in CLOSING or CLOSED state");var e=this,o=h.serverLookup(this.url);if(!o)throw new Error("SocketIO can not find a server at the specified URL ("+this.url+")");return{emit:function(s,a){return o.emit(s,a,{websockets:h.websocketsLookup(e.url,null,e)}),e},to:function(s){return o.to(s,e)},in:function(s){return o.in(s,e)}}},t.prototype.on=function(o,i){return this.addEventListener(o,i),this},t.prototype.off=function(o,i){this.removeEventListener(o,i)},t.prototype.hasListeners=function(o){var i=this.listeners[o];return Array.isArray(i)?!!i.length:!1},t.prototype.join=function(o){h.addMembershipToRoom(this,o)},t.prototype.leave=function(o){h.removeMembershipFromRoom(this,o)},t.prototype.to=function(o){return this.broadcast.to(o)},t.prototype.in=function(){return this.to.apply(null,arguments)},t.prototype.dispatchEvent=function(o){for(var i=this,s=[],a=arguments.length-1;a-- >0;)s[a]=arguments[a+1];var u=o.type,l=this.listeners[u];if(!Array.isArray(l))return!1;l.forEach(function(p){s.length>0?p.apply(i,s):p.call(i,o.data?o.data:o)})},Object.defineProperties(t.prototype,n),t}(C);N.CONNECTING=0,N.OPEN=1,N.CLOSING=2,N.CLOSED=3;var x=function(t,n){return new N(t,n)};x.connect=function(t,n){return x(t,n)};var Re=function(r){return r.reduce(function(t,n){return t.indexOf(n)>-1?t:t.concat(n)},[])};function re(){return typeof window<"u"?window:typeof process=="object"&&typeof Me=="function"&&typeof M=="object"?M:this}var oe={mock:!0,verifyClient:null,selectProtocol:null},z=function(r){function t(n,e){e===void 0&&(e=oe),r.call(this);var o=new B(n);o.pathname||(o.pathname="/"),this.url=o.toString(),this.originalWebSocket=null;var i=h.attachServer(this,this.url);if(!i)throw this.dispatchEvent(E({type:"error"})),new Error("A mock server is already listening on this url");this.options=Object.assign({},oe,e),this.options.mock&&this.mockWebsocket()}return r&&(t.__proto__=r),t.prototype=Object.create(r&&r.prototype),t.prototype.constructor=t,t.prototype.mockWebsocket=function(){var e=re();this.originalWebSocket=e.WebSocket,e.WebSocket=f},t.prototype.restoreWebsocket=function(){var e=re();this.originalWebSocket!==null&&(e.WebSocket=this.originalWebSocket),this.originalWebSocket=null},t.prototype.stop=function(e){e===void 0&&(e=function(){}),this.options.mock&&this.restoreWebsocket(),h.removeServer(this.url),typeof e=="function"&&e()},t.prototype.on=function(e,o){this.addEventListener(e,o)},t.prototype.off=function(e,o){this.removeEventListener(e,o)},t.prototype.close=function(e){e===void 0&&(e={});var o=e.code,i=e.reason,s=e.wasClean,a=h.websocketsLookup(this.url);h.removeServer(this.url),a.forEach(function(u){u.readyState=f.CLOSED,u.dispatchEvent(S({type:"close",target:u.target,code:o||m.CLOSE_NORMAL,reason:i||"",wasClean:s}))}),this.dispatchEvent(S({type:"close"}),this)},t.prototype.emit=function(e,o,i){var s=this;i===void 0&&(i={});var a=i.websockets;a||(a=h.websocketsLookup(this.url));var u;typeof i!="object"||arguments.length>3?(o=Array.prototype.slice.call(arguments,1,arguments.length),u=o.map(function(l){return D(l)})):u=D(o),a.forEach(function(l){var p=l instanceof N?o:u;Array.isArray(p)?l.dispatchEvent.apply(l,[A({type:e,data:p,origin:s.url,target:l.target})].concat(p)):l.dispatchEvent(A({type:e,data:p,origin:s.url,target:l.target}))})},t.prototype.clients=function(){return h.websocketsLookup(this.url)},t.prototype.to=function(e,o,i){var s=this;i===void 0&&(i=[]);var a=this,u=Re(i.concat(h.websocketsLookup(this.url,e,o)));return{to:function(l,p){return s.to.call(s,l,p,u)},emit:function(p,c){a.emit(p,c,{websockets:u})}}},t.prototype.in=function(){for(var e=[],o=arguments.length;o--;)e[o]=arguments[o];return this.to.apply(null,e)},t.prototype.simulate=function(e){var o=h.websocketsLookup(this.url);e==="error"&&o.forEach(function(i){i.readyState=f.CLOSED,i.dispatchEvent(E({type:"error",target:i.target}))})},t}(C);z.of=function(t){return new z(t)};var Ce=z,ke=f,Ne=x;R.Server=Ce,R.WebSocket=ke,R.SocketIO=Ne,Object.defineProperty(R,"__esModule",{value:!0})})})(We,I);Te.WebSocket=I.WebSocket;Ae.enabled=!0;const Ie=`ws://${_e.cableUrl}`,Ue=new I.Server(Ie);function W(v,...L){console.log(`WSMock: ${v}`,...L)}function ne(v){return v[Math.floor(Math.random()*v.length)]}Ue.on("connection",v=>{function L(y,_=null,G={}){_&&(y.identifier=JSON.stringify({channel:_,...G})),y.type!="ping"&&W("send msg",y),v.send(JSON.stringify(y))}W("CONNECTED"),v.on("message",y=>{W("MESSAGE FROM WS",y)}),v.on("close",()=>{W("CLOSE FROM WS")}),v.on("error",y=>{W("ERROR FROM WS",y)}),setInterval(()=>{L({type:"ping"})},3e3),setTimeout(()=>{L({type:"confirm_subscription"},"GroupResourceChannel",{groupResource:"deckhouse.io/moduleconfigs"})},500);function R(){Y.spec.settings.update.mode=ne(["Auto","Manual"]),Y.spec.settings.releaseChannel=ne(["Stable","Alpha","Beta"]),L({message:{message_type:"update",message:Y}},"GroupResourceChannel",{groupResource:"deckhouse.io/moduleconfigs"})}setInterval(R,5e3)});export{Ue as mockServer};
