webpackJsonp([2],{1260:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=n(1),a=n(62),r=n(304),i=n(1393),c=n(1398);n.d(e,"HomeModule",(function(){return s}));var s=(function(){function t(){}return t=__decorate([n.i(o.NgModule)({imports:[a.CommonModule,r.a,c.a],declarations:[i.a],providers:[]}),__metadata("design:paramtypes",[])],t)})()},1393:function(t,e,n){"use strict";var o=n(1),a=n(97);n.d(e,"a",(function(){return r}));var r=(function(){function t(t){this._state=t}return t.prototype.ngOnInit=function(){localStorage.removeItem("breadCrumbs"),this._state.notifyChanged("breadCrumbs")},t=__decorate([n.i(o.Component)({selector:"home",styles:[n(1401)],template:n(1406)}),__metadata("design:paramtypes",[a.a])],t)})()},1398:function(t,e,n){"use strict";var o=n(35),a=n(1393);n.d(e,"a",(function(){return i}));var r=[{path:"",component:a.a,children:[]}],i=o.a.forChild(r)},1401:function(t,e){t.exports=""},1406:function(t,e){t.exports='<div class="row">\n  <div class="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-xs-12">\n    home page\n  </div>\n</div>'}});