import unittest

from bson import ObjectId
import utils
from proto.plot_pb2 import ReadPlotRequest, WritePlotRequest
from proto.plot_pb2_grpc import PlotServiceStub


class TestPlot(utils.TestMedia):
    stub: PlotServiceStub

    def setUp(self) -> None:
        super().setUp('plot', PlotServiceStub)

    def test_write_plot(self) -> None:
        plot = 'some random plot string'
        req = WritePlotRequest(plot=plot)
        resp = self.stub.WritePlot(req)
        oid = ObjectId(resp.plot_id)

        expect = {
            '_id': oid,
            'plot': plot
        }
        actual = self.db['plot'].find_one()
        self.assertEqual(expect, actual)

    def test_read_plot(self) -> None:
        plot = 'some random plot string'
        oid: ObjectId = self.db['plot'].insert_one({'plot': plot}).inserted_id

        req = ReadPlotRequest(plot_id=str(oid))
        resp = self.stub.ReadPlot(req)

        expect = plot
        actual = resp.plot
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
