'use strict';

var docReady = require('doc-ready'),
    request = require('then-request'),
    _ = require('lodash'),
    Packery = require('packery'),

    content,

    buildsTmpl = _.template(document.querySelector('#js-tmpl-builds').innerHTML),

    urls = {},
    builds = [],
    buildStatus = {
        success: 'SUCCESS',
        failure: 'FAILURE'
    },
    refreshTimeout,

    classBuildFailure = 'failure';


function initializeUrls() {
    content = document.querySelector('#js-content');

    urls.builds = content.dataset.urlBuilds;
    urls.buildsStatus = content.dataset.urlBuildsStatus;
}

function initializeMonitoring() {
    request('GET', urls.builds).done(function(response) {
        response = JSON.parse(response.getBody());

        refreshTimeout = response.refreshTimeout * 1000;

        initializeLayouts(response.buildsLayout);
        initializeBuilds(response.builds);
    });
}

function initializeLayouts(layoutsList) {
    content.innerHTML = buildsTmpl({
        layouts: layoutsList
    });

    _.forEach(document.querySelectorAll('.js-layout'), function(container) {
        new Packery(container, {});
    });
}

function initializeBuilds(buildsList) {
    getBuildsStatus();
}

function getBuildsStatus() {
    request('GET', urls.buildsStatus).done(function(response) {
        response = JSON.parse(response.getBody());

        var diff = _.filter(response, function(build) {
            return !_.findWhere(builds, build);
        });
        if (diff) {
            builds = response;
            updateBuilds(diff);
        }

        setTimeout(getBuildsStatus, refreshTimeout);
    });
}

function updateBuilds(builds) {
    _.forEach(builds, function(build) {
        var target = document.querySelector('.' + build.id),
            targetCommiter = target.querySelector('.js-commiter'),
            hasFailureClass = target.classList.contains(classBuildFailure),
            failureClassNeeded = build.status === buildStatus.failure;

        targetCommiter.innerHTML = build.lastCommiter;

        if (!failureClassNeeded && hasFailureClass) {
            target.classList.remove(classBuildFailure);
        } else if (failureClassNeeded && !hasFailureClass) {
            target.classList.add(classBuildFailure);
        }
    });
}


docReady(function() {
    initializeUrls();
    initializeMonitoring();
});
