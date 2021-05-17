(function (root, factory) {if (typeof define === 'function' && define.amd) {define(['exports', 'echarts'], factory);} else if (typeof exports === 'object' && typeof exports.nodeName !== 'string') {factory(exports, require('echarts'));} else {factory({}, root.echarts);}}(this, function (exports, echarts) {var log = function (msg) {if (typeof console !== 'undefined') {console && console.error && console.error(msg);}};if (!echarts) {log('ECharts is not Loaded');return;}if (!echarts.registerMap) {log('ECharts Map is not loaded');return;}echarts.registerMap('武陵源区', {"type":"FeatureCollection","features":[{"type":"Feature","id":"430811","properties":{"name":"武陵源区","cp":[110.550433,29.34573],"childNum":1},"geometry":{"type":"Polygon","coordinates":["@@HIHQ@I@@A@G@A@A@IBGFYJQHCBIFCFCHCBQZAF[R]TSHGA@@@@C@IGGEKAE@GHCLAJCFABE@MAI@GB@BCD@ZAHC^CDA@EFSROLCFQJQBG@C@EDMFOF@@EBC@AE@IEG@@GEKAICE@CBEBAD@FNFHFBF@F@DAFGLEHBDDFBDB@@B@BDDBDDBDBF@DBDBDBDBFBF@DBDBH@F@D@JBJ@F@FBL@JBD@HBJDJBLBHBJBB@D@DBFBBBB@B@@@H@TDRD@BJLHPLHJADEDCHKJCH@XJFLALAR@@DHFDH@F@BC@A@Q@CHAHCFGHIBA@@BCBG@ABMBC@CDIDIFGDCFCFAJ@HJDFDBHHLHLHHDB@F@B@D@BADCBM@IFIRIBAHCNCPCFAV@D@TAH@H@LBHALELIFCDAH@LBN@HCDEBEAGEIIMCC@ACCIGSKEAEAaKIGACCEQGKCMBA@MBK@KAGECCMECI@EBEOLA@ABC@CCCCCAA@E@@AE@E@C@CACAE@EAEAC@E@EBC@EC@@FUDILUBADGVe"],"encodeOffsets":[[113207,29938]]}}],"UTF8Encoding":true});}));