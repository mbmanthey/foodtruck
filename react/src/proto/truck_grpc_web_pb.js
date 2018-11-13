/* eslint-disable */
/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./truck_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.TruckServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.TruckServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.proto.TruckServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.proto.TruckServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Truck,
 *   !proto.proto.Response>}
 */
const methodInfo_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Response,
  /** @param {!proto.proto.Truck} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Response.deserializeBinary
);


/**
 * @param {!proto.proto.Truck} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.TruckServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.TruckService/Create',
      request,
      metadata,
      methodInfo_Create,
      callback);
};


/**
 * @param {!proto.proto.Truck} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Response>}
 *     The XHR Node Readable Stream
 */
proto.proto.TruckServicePromiseClient.prototype.create =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.create(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.GetRequest,
 *   !proto.proto.Response>}
 */
const methodInfo_GetAll = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Response,
  /** @param {!proto.proto.GetRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Response.deserializeBinary
);


/**
 * @param {!proto.proto.GetRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.TruckServiceClient.prototype.getAll =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.TruckService/GetAll',
      request,
      metadata,
      methodInfo_GetAll,
      callback);
};


/**
 * @param {!proto.proto.GetRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Response>}
 *     The XHR Node Readable Stream
 */
proto.proto.TruckServicePromiseClient.prototype.getAll =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getAll(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.proto;

