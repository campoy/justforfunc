# justforfunc 14: a twitter bot and systemd (that runs free on GCP)

Today the goal is to build a Twitter bot that will retweet all tweets containing
[#justforfunc](https://twitter.com/search?q=%23justforfunc&src=typd) from the
[@justforfunc](https://twitter.com/justforfunc) account.

To do this we will use Go, the [anaconda](https://github.com/ChimeraCoder/anaconda) library by
[ChimeraCoder](https://twitter.com/ChimeraCoder), systemd, and ... a free micro instance from
Google Cloud Platform that turns out to be free!

<div style="text-align:center">
    <a href="https://www.youtube.com/watch?v=SQeAKSJH4vw&feature=youtu.be&list=PL6">
        <img src="https://img.youtube.com/vi/SQeAKSJH4vw/0.jpg" alt="justforfunc #14: a twitter bot and systemd (that runs free on GCP)">
        <p>justforfunc #14: a twitter bot and systemd (that runs free on GCP)</p>
    </a>
</div>

References:

- anaconda: https://github.com/ChimeraCoder/anaconda
- systemd: https://www.freedesktop.org/wiki/Software/systemd/
- Google Compute Engine: https://cloud.google.com/compute/
- source code: https://github.com/campoy/justforfunc/tree/master/14-twitterbot