const fs = require('fs')

const src = fs.readFileSync('./todo-backend-js-spec/js/specs.js').toString()
eval(src)

const _ = require('underscore')
const Q = require('q')
const jQuery = require('jquery')
const { JSDOM } = require('jsdom')
const chai = require('chai')
const chaiAsPromised = require('chai-as-promised')

const { window } = new JSDOM()
const $ = jQuery(window)
const { expect } = chai
chai.use(chaiAsPromised)

_, Q, $, expect // Mark these as used

const { TEST_SERVER_URL, TEST_SERVER_PORT } = process.env
defineSpecsFor(TEST_SERVER_URL || `http://localhost:${TEST_SERVER_PORT || 8080}`)
