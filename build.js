var path = require('path'),
    webpack = require('webpack'),

    config = {
        context: path.join(__dirname, 'static/js'),
        entry: './index.js',
        output: {
            path: path.join(__dirname, 'static/assets')
        },
        module: {
            loaders: [{
                test: /packery/,
                loader: 'imports?define=>false&this=>window'
            }]
        },
        plugins: [
            new webpack.optimize.UglifyJsPlugin(),
            new webpack.optimize.DedupePlugin()
        ]
    };

webpack(config).run(function() {});
