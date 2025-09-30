// Snowplow tracker initialization
;(function(p,l,o,w,i,n,g){if(!p[i]){p.GlobalSnowplowNamespace=p.GlobalSnowplowNamespace||[];
p.GlobalSnowplowNamespace.push(i);p[i]=function(){(p[i].q=p[i].q||[]).push(arguments)
};p[i].q=p[i].q||[];n=l.createElement(o);g=l.getElementsByTagName(o)[0];n.async=1;
n.src=w;g.parentNode.insertBefore(n,g)}}(window,document,"script","/home/assets/js/hmpjjmjk.js","snowplow"));

var collector = window.location.host;

snowplow('newTracker', 'sp', collector, {
    appId: 'SnowplowMini',
    contexts: {
        webPage: true,
        performanceTiming: true
    }
});

snowplow('disableActivityTracking');

snowplow('setUserId', 'SnowplowMini-User');
