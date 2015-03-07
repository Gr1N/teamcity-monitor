<!DOCTYPE html>
<html>
<head>
  <title>TeamCity Monitor</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <link href="/static/css/style.css" rel="stylesheet" />
</head>

<body>
  <div id="js-content"
       data-url-builds="{{urlfor "APIController.Builds"}}"
       data-url-builds-status="{{urlfor "APIController.BuildsStatus"}}"
  ></div>

  <script type="text/template" id="js-tmpl-builds">
    <% _.forEach(layouts, function(layout) { %>
      <div class="js-layout container">
        <% _.forEach(layout, function(build) { %>
          <div class="item <%- build %>"></div>
        <% }); %>
      </div>
    <% }); %>
  </script>
  <script type="text/template" id="js-tmpl-build">
    <p class="name"><%- name %></p>
    <p class="status"><%- statusText %></p>
    <p class="commiter"><%- lastCommiter %></p>
  </script>
  <script src="/static/assets/bundle.js"></script>
</body>
</html>
