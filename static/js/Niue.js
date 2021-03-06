(function (root, factory) {
    if (typeof define === 'function' && define.amd) {
        define(['exports', 'echarts'], factory);
    } else if (typeof exports === 'object' && typeof exports.nodeName !== 'string') {
        factory(exports, require('echarts'));
    } else {
        factory({}, root.echarts);
    }
}(this, function (exports, echarts) {
    var log = function (msg) {
        if (typeof console !== 'undefined') {
            console && console.error && console.error(msg);
        }
    };
    if (!echarts) {
        log('ECharts is not Loaded');
        return;
    }
    if (!echarts.registerMap) {
        log('ECharts Map is not loaded');
        return;
    }
    echarts.registerMap('纽埃', {
        "type": "FeatureCollection",
        "features": [{
            "type": "Feature",
            "properties": {"name": "Alofi"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@mNBn\\ZACcC]uuKMMJuNEAKFEEuJBDFFD@\\TF@DFHDBDFN@L@HCJERAN"],
                "encodeOffsets": [[-174001, -19472]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Avatele"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@JBHXvTJKIgAIG@@IE@CEGDM[IDMBIAGBED@DNJFFBHAJAH"],
                "encodeOffsets": [[-173991, -19571]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Hakupu"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@LNvvdkvmGEEGIKMKIEIIKMGO@AGGQ@WEKCIEA@BzEBFD@FCBAHJhIL"],
                "encodeOffsets": [[-173955, -19548]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Hikutavake"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@A]XWW_ghQTDDFBDFHFHDB@L@DBF@FB"],
                "encodeOffsets": [[-173938, -19411]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Lakepa"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@`\\DVZBJVTA@HNBBKB@FSJQBEAGACAIcMcE[UDdYB"],
                "encodeOffsets": [[-173931, -19480]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Liku"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@D^\\VdFdNDIDCBEDEBCDGB@DI@CJMDADIAGEG@CCIOSCCCECEGKIEuncl"],
                "encodeOffsets": [[-173922, -19514]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Makefu"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@TE^Wz]m[MAnAFFNPNFFFD"],
                "encodeOffsets": [[-173985, -19450]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Mutalau"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@X`VZMVBFE@CHL@VCPEL@FALCHCFGDKBI@KAAMA@GSBIUYACU_[Ot"],
                "encodeOffsets": [[-173939, -19454]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Namukulu"},
            "geometry": {"type": "Polygon", "coordinates": ["@@RScEBHHLHF"], "encodeOffsets": [[-173968, -19424]]}
        }, {
            "type": "Feature",
            "properties": {"name": "Tamakautoga"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@vIFFLEFBvMNIuSGWIAELABUTSLEFOTAF"],
                "encodeOffsets": [[-174028, -19532]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Toi"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@B@FBNFDGF@AENUUYWXB^"],
                "encodeOffsets": [[-173938, -19411]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Tuapa"},
            "geometry": {"type": "Polygon", "coordinates": ["@@dFhgPsy^]XSFFHJR"], "encodeOffsets": [[-173977, -19437]]}
        }, {
            "type": "Feature",
            "properties": {"name": "Vaiea"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@N\\HCDFF@@JH@BJBGDA@EECFAAyEAGFOPOJ"],
                "encodeOffsets": [[-173976, -19599]]
            }
        }],
        "UTF8Encoding": true
    });
}));